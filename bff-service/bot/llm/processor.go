package llm

import (
	"context"
	"errors"
	"fmt"
)

type EventType string

const (
	FieldFilled   EventType = "field_filled"
	FieldChanged  EventType = "field_changed"
	FormCompleted EventType = "form_completed"
)

type Event struct {
	Type      EventType
	SessionID string
	FormID    string
	Field     string
	Old, New  string // valores anterior y nuevo (útil para auditoría)
	Answers   Answers
}

// Listener (Observer): reacciona a eventos sin que el motor sepa qué hace
// (guardar en BD, llamar un webhook, métricas, enviar correo...).
type Listener func(ctx context.Context, e Event)

type Processor struct {
	form         Form
	store        Store
	handlers     []Handler
	extractor    FormExtractor
	processFirst bool
	confirm      bool
	listeners    []Listener
}

// Option (Functional Options): configuración extensible sin romper la firma de NewProcessor.
type Option func(*Processor)

func WithStore(s Store) Option         { return func(p *Processor) { p.store = s } }
func WithHandlers(h ...Handler) Option { return func(p *Processor) { p.handlers = h } }
func OnEvent(l Listener) Option        { return func(p *Processor) { p.listeners = append(p.listeners, l) } }

// WithFormExtractor cambia la estrategia de extracción (reglas por campo o LLM multi-campo).
func WithFormExtractor(fe FormExtractor) Option { return func(p *Processor) { p.extractor = fe } }

// ProcessFirstMessage hace que el primer mensaje de una sesión también se procese
// (útil con LLM: "Hola, soy Ana, mi correo es ana@x.com"). Por defecto se ignora
// y solo se responde con la primera pregunta.
func ProcessFirstMessage() Option { return func(p *Processor) { p.processFirst = true } }

// RequireConfirmation agrega una fase final: se muestra un resumen y el usuario
// confirma ("sí") o elige qué dato corregir ("no").
func RequireConfirmation() Option { return func(p *Processor) { p.confirm = true } }

func NewProcessor(form Form, opts ...Option) *Processor {
	p := &Processor{
		form:      form,
		store:     NewMemoryStore(),
		handlers:  DefaultHandlers(),
		extractor: PerFieldExtractor{},
	}
	for _, o := range opts {
		o(p)
	}
	return p
}

func (p *Processor) Handle(ctx context.Context, sessionID, input string) (Reply, error) {
	first := false
	s, err := p.store.Load(ctx, sessionID)
	switch {
	case errors.Is(err, ErrSessionNotFound):
		if _, ok := p.form.Next(Answers{}); !ok {
			return Reply{}, errors.New("el formulario no tiene campos")
		}
		s = NewSession(sessionID, p.form.ID)
		if !p.processFirst {
			if err := p.store.Save(ctx, s); err != nil {
				return Reply{}, fmt.Errorf("guardando sesión: %w", err)
			}
			f, _ := p.form.Next(s.Answers)
			return Reply{Text: f.Prompt}, nil
		}
		first = true
	case err != nil:
		return Reply{}, fmt.Errorf("cargando sesión: %w", err)
	case s.Done:
		return Reply{Text: "Esta conversación ya terminó.", Done: true}, nil
	}

	turn := &Turn{
		Session:   s,
		Form:      p.form,
		Input:     input,
		Extractor: p.extractor,
		First:     first,
		Confirm:   p.confirm,
		Emit: func(e Event) {
			e.SessionID, e.FormID = s.ID, s.FormID
			for _, l := range p.listeners {
				l(ctx, e)
			}
		},
	}

	for _, h := range p.handlers {
		reply, handled, err := h.Handle(ctx, turn)
		if err != nil {
			return Reply{}, err // no se guarda: la sesión queda como estaba
		}
		if !handled {
			continue
		}
		if err := p.store.Save(ctx, s); err != nil {
			return Reply{}, fmt.Errorf("guardando sesión: %w", err)
		}
		return reply, nil
	}
	return Reply{}, errors.New("ningún handler procesó el mensaje")
}
