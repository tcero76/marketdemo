package llm

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/tcero76/marketplace/bff-service/services"
	modelServices "github.com/tcero76/marketplace/postgres/services"
	redisServices "github.com/tcero76/marketplace/redis/services"

	"github.com/tcero76/marketplace/redis/model"
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
	Answers   model.Answers
}

// Listener (Observer): reacciona a eventos sin que el motor sepa qué hace
// (guardar en BD, llamar un webhook, métricas, enviar correo...).
type Listener func(ctx context.Context, e Event)

type Processor struct {
	form         Form
	store        services.IAuthCacheService
	handlers     []Handler
	extractor    FormExtractor
	ProcessFirst bool
	confirm      bool
	listeners    []Listener
}

// Option (Functional Options): configuración extensible sin romper la firma de NewProcessor.
type Option func(*Processor)

func OnEvent(l Listener) Option { return func(p *Processor) { p.listeners = append(p.listeners, l) } }

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// ---- Validadores propios del dominio ----

func futureDate(today time.Time) Validator {
	midnight := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)
	return ValidatorFunc(func(v string) error {
		d, err := time.Parse("2006-01-02", v)
		if err != nil {
			return Invalid("Usa el formato AAAA-MM-DD")
		}
		if d.Before(midnight) {
			return Invalid("La fecha no puede estar en el pasado")
		}
		return nil
	})
}

func between(min, max int) Validator {
	return ValidatorFunc(func(v string) error {
		n, err := strconv.Atoi(v)
		if err != nil || n < min || n > max {
			return Invalid("Debe ser un número entre %d y %d", min, max)
		}
		return nil
	})
}

// ---- Observer: persiste cada reserva completada ----

func saveReservation(chatbotServices *modelServices.Service) Listener {
	return func(ctx context.Context, e Event) {
		if e.Type != FormCompleted {
			return
		}
		answers := make(map[string]string, len(e.Answers))
		for key, value := range e.Answers {
			answers[key] = value
		}
		chatbotServices.Create(ctx, answers, e.SessionID)
	}
}

func NewBotLLM(authCacheService *redisServices.AuthCacheService, chatbotServices *modelServices.Service) *Processor {
	today := time.Now()
	form := Form{ID: "reserva", Fields: []Field{
		{Name: "nombre", Label: "Nombre", Prompt: "¿A nombre de quién es la reserva?",
			Hint:       "nombre de la persona que reserva",
			Validators: []Validator{MinLength(2)}},
		{Name: "email", Label: "Correo", Prompt: "¿Cuál es tu correo?",
			Hint:       "correo electrónico",
			Validators: []Validator{Matches(EmailRe, "Correo inválido")}},
		{Name: "fecha", Label: "Fecha", Prompt: "¿Para qué fecha? (AAAA-MM-DD)",
			Hint:       "fecha de la reserva en formato AAAA-MM-DD; hoy es " + today.Format("2006-01-02"),
			Validators: []Validator{futureDate(today)}},
		{Name: "personas", Label: "Personas", Prompt: "¿Para cuántas personas? (1-10)",
			Hint:       "cantidad de personas, solo el número entero",
			Type:       "integer",
			Validators: []Validator{between(1, 10)}},
	}}
	client := OpenAICompatClient{
		BaseURL: getenv("LLM_URL", "http://localhost:12434/engines/v1"),
		Model:   getenv("LLM_MODEL", "qwen2.5"),
	}
	debug := os.Getenv("DEBUG") == "1"
	var extractor FormExtractor
	if getenv("EXTRACTOR", "tools") == "prompt" {
		var c Completer = client
		if debug {
			c = loggingCompleter{next: c}
		}
		extractor = LLMFormExtractor{Client: c}
	} else {
		var c ToolCaller = client
		if debug {
			c = loggingToolCaller{next: c}
		}
		extractor = LLMToolExtractor{Client: c}
	}
	p := &Processor{
		form:         form,
		store:        authCacheService,
		handlers:     DefaultHandlers(),
		extractor:    extractor,
		confirm:      true,
		ProcessFirst: true,
		listeners:    []Listener{saveReservation(chatbotServices)},
	}
	return p
}

func (p *Processor) Handle(ctx context.Context, sessionID, input string, s *model.FormSession, first bool) (Reply, error) {
	fmt.Println("Entró al Handler")
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
		if err := p.store.StoreTokenInRedis(s.ID, "bot_session", s, ctx); err != nil {
			return Reply{}, fmt.Errorf("guardando sesión: %w", err)
		}
		return reply, nil
	}
	return Reply{}, errors.New("ningún handler procesó el mensaje")
}
