package llm

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

// Turn es el contexto de un mensaje entrante que viaja por la cadena.
type Turn struct {
	Session   *Session
	Form      Form
	Input     string
	Extractor FormExtractor // estrategia de extracción configurada en el Processor
	First     bool          // true si es el primer mensaje de la sesión
	Confirm   bool          // true si el formulario exige confirmación final
	Emit      func(Event)
}

type Reply struct {
	Text    string
	Done    bool    // la conversación terminó (completa o cancelada)
	Answers Answers // solo se llena al completar
}

// Handler (Chain of Responsibility): cada eslabón decide si atiende el mensaje.
// handled=false pasa el turno al siguiente.
type Handler interface {
	Handle(ctx context.Context, t *Turn) (reply Reply, handled bool, err error)
}

// DefaultHandlers define el orden de la cadena. El ORDEN es parte del diseño:
//
//  1. Cancel / Back          comandos exactos, valen en cualquier fase.
//  2. Confirm (estricto)     "sí" / "no" cortos durante la confirmación.
//  3. Fill                   llena campos pendientes y REVISA los ya llenos.
//  4. Confirm (flexible)     "sí, está bien…" si Fill no encontró cambios.
//  5. PickField              el usuario nombra qué dato corregir.
func DefaultHandlers() []Handler {
	yes := []string{"si", "sí", "s", "ok", "vale", "correcto", "confirmo", "yes"}
	no := []string{"no", "n"}
	return []Handler{
		CancelHandler{Keywords: []string{"cancelar", "cancel"}},
		BackHandler{Keywords: []string{"volver", "atras", "atrás", "back"}},
		ConfirmHandler{Yes: yes, No: no},
		FillHandler{},
		ConfirmHandler{Yes: yes, No: no, Lenient: true},
		PickFieldHandler{},
	}
}

// ---------- utilidades ----------

func matchesAny(input string, kws []string) bool {
	in := strings.ToLower(strings.TrimSpace(input))
	for _, k := range kws {
		if in == strings.ToLower(k) {
			return true
		}
	}
	return false
}

func contains(list []string, v string) bool {
	for _, x := range list {
		if x == v {
			return true
		}
	}
	return false
}

func nextPrompt(t *Turn) string {
	if f, ok := t.Form.Next(t.Session.Answers); ok {
		return f.Prompt
	}
	return ""
}

func fieldLabels(t *Turn) string {
	var names []string
	for _, f := range t.Form.Fields {
		if _, ok := t.Session.Answers[f.Name]; ok {
			names = append(names, f.DisplayName())
		}
	}
	return strings.Join(names, ", ")
}

func summary(t *Turn) string {
	var b strings.Builder
	b.WriteString("Esto es lo que tengo:\n")
	for _, f := range t.Form.Fields {
		if v, ok := t.Session.Answers[f.Name]; ok {
			fmt.Fprintf(&b, "- %s: %s\n", f.DisplayName(), v)
		}
	}
	b.WriteString("¿Es correcto? (sí / no)")
	return b.String()
}

func finish(t *Turn) Reply {
	s := t.Session
	s.Done = true
	answers := s.clone().Answers
	t.Emit(Event{Type: FormCompleted, Answers: answers})
	return Reply{Text: "¡Listo! Ya tengo todos tus datos.", Done: true, Answers: answers}
}

func removeString(list []string, v string) []string {
	var out []string
	for _, x := range list {
		if x != v {
			out = append(out, x)
		}
	}
	return out
}

// findField busca, entre los campos ya respondidos, el que el usuario menciona.
func findField(form Form, answers Answers, input string) (Field, bool) {
	low := strings.ToLower(input)
	for _, f := range form.Fields {
		if _, ok := answers[f.Name]; !ok {
			continue
		}
		if strings.Contains(low, strings.ToLower(f.Name)) {
			return f, true
		}
		if f.Label != "" && strings.Contains(low, strings.ToLower(f.Label)) {
			return f, true
		}
	}
	return Field{}, false
}

// ---------- Cancelar ----------

type CancelHandler struct{ Keywords []string }

func (h CancelHandler) Handle(_ context.Context, t *Turn) (Reply, bool, error) {
	if !matchesAny(t.Input, h.Keywords) {
		return Reply{}, false, nil
	}
	t.Session.Done = true
	return Reply{Text: "Formulario cancelado.", Done: true}, true, nil
}

// ---------- Volver al paso anterior ----------

type BackHandler struct{ Keywords []string }

func (h BackHandler) Handle(_ context.Context, t *Turn) (Reply, bool, error) {
	if !matchesAny(t.Input, h.Keywords) {
		return Reply{}, false, nil
	}
	s := t.Session
	if n := len(s.History); n > 0 {
		delete(s.Answers, s.History[n-1])
		s.History = s.History[:n-1]
	}
	s.Phase = PhaseCollecting
	return Reply{Text: nextPrompt(t)}, true, nil
}

// ---------- Confirmación (sí / no) ----------

// ConfirmHandler responde al resumen final. En modo estricto solo acepta mensajes
// cortos ("sí", "no", "sí, correcto"): así "sí, pero a nombre de Paula" NO se toma
// como un "sí" y llega a FillHandler como revisión. En modo flexible (después de
// FillHandler) acepta "sí, está bien, gracias" mirando la primera palabra, y
// responde cuando no entiende.
type ConfirmHandler struct {
	Yes, No []string
	Lenient bool
}

func (h ConfirmHandler) Handle(_ context.Context, t *Turn) (Reply, bool, error) {
	s := t.Session
	if s.Phase != PhaseConfirming {
		return Reply{}, false, nil
	}
	yes, no := h.classify(t.Input)
	switch {
	case yes:
		return finish(t), true, nil
	case no:
		s.Phase = PhaseCorrecting
		return Reply{Text: "¿Qué dato quieres corregir? (" + fieldLabels(t) + ")"}, true, nil
	case h.Lenient:
		return Reply{Text: "Responde «sí» para confirmar o «no» para corregir un dato."}, true, nil
	}
	return Reply{}, false, nil
}

func (h ConfirmHandler) classify(input string) (yes, no bool) {
	words := strings.Fields(strings.ToLower(input))
	if len(words) == 0 || (!h.Lenient && len(words) > 2) {
		return false, false
	}
	w := strings.Trim(words[0], ",.;:!¡¿?")
	return contains(h.Yes, w), contains(h.No, w)
}

// ---------- Llenar campos y revisar los ya llenos ----------

type fieldChange struct{ name, label, old, new string }

// FillHandler aplica lo que el FormExtractor encontró:
//   - campos PENDIENTES  → se llenan (si pasan las validaciones);
//   - campos YA LLENOS   → se cambian solo si el valor es distinto (revisión);
//   - llaves desconocidas → se ignoran (protege contra campos inventados).
//
// Tras un cambio se limpian las respuestas de campos que Skip ahora omite
// (p. ej. cambiar canal de "telefono" a "email" descarta el teléfono).
type FillHandler struct{}

func (FillHandler) Handle(ctx context.Context, t *Turn) (Reply, bool, error) {
	s := t.Session
	pending := t.Form.Pending(s.Answers)
	if len(pending) == 0 && s.Phase == PhaseCollecting {
		return Reply{}, false, nil
	}

	fe := t.Extractor
	if fe == nil {
		fe = PerFieldExtractor{}
	}
	candidates, err := fe.ExtractAll(ctx, ExtractRequest{
		Pending: pending,
		Filled:  t.Form.Filled(s.Answers),
		Answers: s.Answers,
		Input:   t.Input,
	})
	if err != nil && !errors.Is(err, ErrNotUnderstood) {
		return Reply{}, true, err // error de sistema
	}

	var filled, filledLabels, problems []string
	var changed []fieldChange
	for _, f := range t.Form.Fields {
		raw, ok := candidates[f.Name]
		if !ok {
			continue
		}
		if f.Skip != nil && f.Skip(s.Answers) {
			continue
		}
		value := strings.TrimSpace(raw)
		if value == "" {
			continue
		}
		old, had := s.Answers[f.Name]
		if had && old == value {
			continue // el LLM repitió el valor actual: no es un cambio
		}
		if err := runValidators(f.AllValidators(), value); err != nil {
			var ve *ValidationError
			if errors.As(err, &ve) {
				problems = append(problems, fmt.Sprintf("%s: %s.", f.DisplayName(), ve.Msg))
				continue
			}
			return Reply{}, true, err
		}
		s.Answers[f.Name] = value
		if had {
			changed = append(changed, fieldChange{f.Name, f.DisplayName(), old, value})
		} else {
			s.History = append(s.History, f.Name)
			filled = append(filled, f.Name)
			filledLabels = append(filledLabels, f.DisplayName())
		}
	}

	if len(changed) > 0 {
		for _, f := range t.Form.Fields {
			if _, ok := s.Answers[f.Name]; ok && f.Skip != nil && f.Skip(s.Answers) {
				delete(s.Answers, f.Name)
				s.History = removeString(s.History, f.Name)
			}
		}
	}

	nothing := len(filled) == 0 && len(changed) == 0 && len(problems) == 0
	if nothing && s.Phase != PhaseCollecting {
		return Reply{}, false, nil // en confirmación/corrección: que lo atiendan los siguientes
	}

	for _, name := range filled {
		t.Emit(Event{Type: FieldFilled, Field: name, New: s.Answers[name]})
	}
	for _, c := range changed {
		t.Emit(Event{Type: FieldChanged, Field: c.name, Old: c.old, New: c.new})
	}

	var parts []string
	if len(filled) > 1 {
		parts = append(parts, "Registré: "+strings.Join(filledLabels, ", ")+".")
	}
	for _, c := range changed {
		parts = append(parts, fmt.Sprintf("Cambié %s: %s → %s.", c.label, c.old, c.new))
	}
	parts = append(parts, problems...)

	next, more := t.Form.Next(s.Answers)
	if !more {
		if t.Confirm {
			s.Phase = PhaseConfirming
			parts = append(parts, summary(t))
			return Reply{Text: strings.Join(parts, " ")}, true, nil
		}
		return finish(t), true, nil
	}

	s.Phase = PhaseCollecting // p. ej. un cambio hizo aparecer un campo nuevo
	if nothing && !t.First {
		parts = append(parts, ErrNotUnderstood.Error()+".")
	}
	parts = append(parts, next.Prompt)
	return Reply{Text: strings.Join(parts, " ")}, true, nil
}

// ---------- Elegir qué dato corregir (por nombre) ----------

// PickFieldHandler es el último recurso en la fase de corrección: el usuario dijo
// "no" y ahora nombra el campo ("la fecha"). Si además daba un valor nuevo, ya lo
// habría atendido FillHandler como revisión.
type PickFieldHandler struct{}

func (PickFieldHandler) Handle(_ context.Context, t *Turn) (Reply, bool, error) {
	s := t.Session
	if s.Phase != PhaseCorrecting {
		return Reply{}, false, nil
	}
	f, ok := findField(t.Form, s.Answers, t.Input)
	if !ok {
		return Reply{Text: "No reconocí ese dato. Opciones: " + fieldLabels(t) + "."}, true, nil
	}
	delete(s.Answers, f.Name)
	s.History = removeString(s.History, f.Name)
	s.Phase = PhaseCollecting
	return Reply{Text: f.Prompt}, true, nil
}
