package llm

import (
	"context"
	"strings"
	"testing"
)

func testForm() Form {
	return Form{ID: "contacto", Fields: []Field{
		{Name: "nombre", Prompt: "¿Cómo te llamas?", Validators: []Validator{MinLength(2)}},
		{Name: "email", Prompt: "¿Cuál es tu correo?",
			Extractor:  RegexExtractor{Re: EmailRe},
			Validators: []Validator{Matches(EmailRe, "Correo inválido")}},
		{Name: "canal", Prompt: "¿Prefieres email o telefono?",
			Extractor:  ChoiceExtractor{Options: []string{"email", "telefono"}},
			Enum:       []string{"email", "telefono"}},
		{Name: "telefono", Prompt: "¿Cuál es tu teléfono?",
			Skip: func(a Answers) bool { return a["canal"] == "email" }},
	}}
}

func say(t *testing.T, p *Processor, in string) Reply {
	t.Helper()
	r, err := p.Handle(context.Background(), "s1", in)
	if err != nil {
		t.Fatalf("Handle(%q): %v", in, err)
	}
	return r
}

type fakeLLM struct{ out string }

func (f fakeLLM) Complete(context.Context, string) (string, error) { return f.out, nil }

// ---- modo reglas (un campo por turno) ----

func TestRulesHappyPathWithSkipAndEvents(t *testing.T) {
	var events []EventType
	p := NewProcessor(testForm(), OnEvent(func(_ context.Context, e Event) { events = append(events, e.Type) }))

	if r := say(t, p, "hola"); r.Text != "¿Cómo te llamas?" {
		t.Fatalf("primera pregunta inesperada: %q", r.Text)
	}
	say(t, p, "Ana")
	if r := say(t, p, "no sé"); !strings.Contains(r.Text, "correo") {
		t.Fatalf("debía re-preguntar el correo: %q", r.Text)
	}
	say(t, p, "mi correo es ana@x.com")
	r := say(t, p, "email") // "telefono" se omite por Skip

	if !r.Done || r.Answers["email"] != "ana@x.com" || r.Answers["nombre"] != "Ana" {
		t.Fatalf("respuesta final inesperada: %+v", r)
	}
	if _, ok := r.Answers["telefono"]; ok {
		t.Fatal("telefono debía omitirse")
	}
	if events[len(events)-1] != FormCompleted {
		t.Fatalf("faltó el evento de cierre: %v", events)
	}
}

func TestBackAndCancel(t *testing.T) {
	p := NewProcessor(testForm())
	say(t, p, "hola")
	say(t, p, "Ana")
	if r := say(t, p, "volver"); r.Text != "¿Cómo te llamas?" {
		t.Fatalf("volver debía repetir la pregunta anterior: %q", r.Text)
	}
	if r := say(t, p, "cancelar"); !r.Done {
		t.Fatal("cancelar debía terminar")
	}
	if r := say(t, p, "hola?"); !strings.Contains(r.Text, "terminó") {
		t.Fatalf("sesión terminada debía avisarlo: %q", r.Text)
	}
}

// ---- modo LLM (varios campos por mensaje) ----

func TestLLMFillsManyFieldsFromFirstMessage(t *testing.T) {
	llm := fakeLLM{out: "```json\n{\"nombre\":\"Ana\",\"email\":\"ana@x.com\",\"canal\":\"email\",\"inventado\":\"x\"}\n```"}
	p := NewProcessor(testForm(), WithFormExtractor(LLMFormExtractor{Client: llm}), ProcessFirstMessage())

	r := say(t, p, "Hola, soy Ana, ana@x.com, prefiero email")
	if !r.Done || r.Answers["nombre"] != "Ana" || r.Answers["canal"] != "email" {
		t.Fatalf("debía completarse en un solo mensaje: %+v", r)
	}
	if _, ok := r.Answers["inventado"]; ok {
		t.Fatal("un campo inexistente no debe aceptarse")
	}
}

func TestLLMPartialAndInvalid(t *testing.T) {
	llm := fakeLLM{out: `{"nombre":"A","email":"ana@x.com"}`} // nombre muy corto
	p := NewProcessor(testForm(), WithFormExtractor(LLMFormExtractor{Client: llm}), ProcessFirstMessage())

	r := say(t, p, "soy A, ana@x.com")
	if r.Done || !strings.Contains(r.Text, "nombre:") || !strings.Contains(r.Text, "¿Cómo te llamas?") {
		t.Fatalf("debía reportar el nombre inválido y re-preguntarlo: %q", r.Text)
	}
}

func TestLLMGreetingAsFirstMessage(t *testing.T) {
	p := NewProcessor(testForm(), WithFormExtractor(LLMFormExtractor{Client: fakeLLM{out: "{}"}}), ProcessFirstMessage())
	if r := say(t, p, "hola"); r.Text != "¿Cómo te llamas?" {
		t.Fatalf("un saludo debía responderse solo con la primera pregunta: %q", r.Text)
	}
}

// ---- confirmación y corrección ----

func TestConfirmationAndCorrection(t *testing.T) {
	p := NewProcessor(testForm(), RequireConfirmation())
	say(t, p, "hola")
	say(t, p, "Ana")
	say(t, p, "ana@x.com")
	r := say(t, p, "email") // canal; telefono se omite

	if r.Done || !strings.Contains(r.Text, "¿Es correcto?") || !strings.Contains(r.Text, "ana@x.com") {
		t.Fatalf("debía pedir confirmación con resumen: %+v", r)
	}
	if r := say(t, p, "no"); !strings.Contains(r.Text, "¿Qué dato") {
		t.Fatalf("debía preguntar qué corregir: %q", r.Text)
	}
	if r := say(t, p, "email"); r.Text != "¿Cuál es tu correo?" {
		t.Fatalf("debía re-preguntar el correo: %q", r.Text)
	}
	r = say(t, p, "otro@x.com")
	if r.Done || !strings.Contains(r.Text, "otro@x.com") {
		t.Fatalf("debía volver al resumen con el dato corregido: %+v", r)
	}
	r = say(t, p, "sí, correcto")
	if !r.Done || r.Answers["email"] != "otro@x.com" {
		t.Fatalf("debía completarse con el correo corregido: %+v", r)
	}
}

// ---- revisiones en lenguaje natural ----

// seqLLM devuelve respuestas predefinidas, una por llamada.
type seqLLM struct {
	outs []string
	i    int
}

func (s *seqLLM) Complete(context.Context, string) (string, error) {
	if s.i >= len(s.outs) {
		return "{}", nil
	}
	out := s.outs[s.i]
	s.i++
	return out, nil
}

func llmProcessor(outs ...string) *Processor {
	return NewProcessor(testForm(),
		WithFormExtractor(LLMFormExtractor{Client: &seqLLM{outs: outs}}),
		ProcessFirstMessage(),
		RequireConfirmation(),
	)
}

func TestRevisionNaturalLanguage(t *testing.T) {
	p := llmProcessor(
		`{"nombre":"Ana","email":"ana@x.com","canal":"email"}`,
		`{"nombre":"Paula"}`,
	)
	r := say(t, p, "Soy Ana, ana@x.com, prefiero email")
	if r.Done || !strings.Contains(r.Text, "¿Es correcto?") {
		t.Fatalf("debía pedir confirmación: %+v", r)
	}

	r = say(t, p, "Cambié de opinión, deja la reserva del restaurant a nombre de Paula")
	if !strings.Contains(r.Text, "Ana → Paula") || !strings.Contains(r.Text, "¿Es correcto?") {
		t.Fatalf("debía aplicar el cambio y volver al resumen: %q", r.Text)
	}

	r = say(t, p, "sí")
	if !r.Done || r.Answers["nombre"] != "Paula" {
		t.Fatalf("debía completarse con el nombre nuevo: %+v", r)
	}
}

func TestRevisionClearsDependentFields(t *testing.T) {
	p := llmProcessor(
		`{"nombre":"Ana","email":"ana@x.com","canal":"telefono"}`,
		`{"telefono":"555-1234"}`,
		`{"canal":"email"}`,
	)
	r := say(t, p, "Soy Ana, ana@x.com, llámame")
	if !strings.Contains(r.Text, "¿Cuál es tu teléfono?") {
		t.Fatalf("debía pedir el teléfono: %q", r.Text)
	}
	r = say(t, p, "es el 555-1234")
	if !strings.Contains(r.Text, "555-1234") || !strings.Contains(r.Text, "¿Es correcto?") {
		t.Fatalf("debía mostrar el resumen con teléfono: %q", r.Text)
	}
	r = say(t, p, "mejor por email, ya no me llamen")
	if !strings.Contains(r.Text, "¿Es correcto?") || strings.Contains(r.Text, "555-1234") {
		t.Fatalf("el teléfono debía descartarse al cambiar a email: %q", r.Text)
	}
}

func TestNoopRevisionIsIgnored(t *testing.T) {
	p := llmProcessor(
		`{"nombre":"Ana","email":"ana@x.com","canal":"email"}`,
		`{"nombre":"Ana"}`, // el LLM repite el valor actual
	)
	say(t, p, "Soy Ana, ana@x.com, prefiero email")
	r := say(t, p, "mmm no sé qué decir")
	if !strings.Contains(r.Text, "Responde «sí»") {
		t.Fatalf("un 'cambio' al mismo valor no debe contar: %q", r.Text)
	}
}

// ---- function calling (tools) ----

type fakeTool struct {
	args string
	got  ToolSpec
}

func (f *fakeTool) CallTool(_ context.Context, _, _ string, tool ToolSpec) (string, error) {
	f.got = tool
	return f.args, nil
}

func TestToolSchemaFromForm(t *testing.T) {
	form := Form{ID: "x", Fields: []Field{
		{Name: "nombre", Prompt: "¿Nombre?", Hint: "nombre de quien reserva"},
		{Name: "personas", Prompt: "¿Cuántas?", Type: "integer"},
		{Name: "canal", Prompt: "¿Canal?", Enum: []string{"email", "telefono"}},
	}}
	tool := &fakeTool{args: `{"nombre":"Ana"}`}
	p := NewProcessor(form, WithFormExtractor(LLMToolExtractor{Client: tool}), ProcessFirstMessage())
	say(t, p, "soy Ana")

	props, ok := tool.got.Parameters["properties"].(map[string]any)
	if !ok {
		t.Fatalf("el schema debía tener properties: %+v", tool.got.Parameters)
	}
	if props["personas"].(map[string]any)["type"] != "integer" {
		t.Fatalf("personas debía ser integer: %+v", props["personas"])
	}
	if enum, _ := props["canal"].(map[string]any)["enum"].([]string); len(enum) != 2 {
		t.Fatalf("canal debía llevar enum: %+v", props["canal"])
	}
	if desc, _ := props["nombre"].(map[string]any)["description"].(string); !strings.Contains(desc, "nombre de quien reserva") {
		t.Fatalf("la descripción debía usar el Hint: %q", desc)
	}
}

func TestToolExtractorEnumIsEnforced(t *testing.T) {
	tool := &fakeTool{args: `{"canal":"sms"}`} // fuera del enum
	p := NewProcessor(testForm(), WithFormExtractor(LLMToolExtractor{Client: tool}), ProcessFirstMessage())
	r := say(t, p, "por sms")
	if !strings.Contains(r.Text, "canal:") {
		t.Fatalf("un valor fuera del enum debía rechazarse: %q", r.Text)
	}
}

func TestToolArgumentsWrappedAsText(t *testing.T) {
	tool := &fakeTool{args: "<tool_call>\n{\"name\": \"registrar_datos\", \"arguments\": {\"nombre\": \"Ana\"}}\n</tool_call>"}
	p := NewProcessor(testForm(), WithFormExtractor(LLMToolExtractor{Client: tool}), ProcessFirstMessage())
	r := say(t, p, "soy Ana")
	if r.Text != "¿Cuál es tu correo?" {
		t.Fatalf("debía registrar el nombre y pedir el correo: %q", r.Text)
	}
}
