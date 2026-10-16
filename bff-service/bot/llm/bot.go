package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// ---- Decorator: agrega logs a cualquier Completer sin tocar el original ----

type loggingCompleter struct{ next Completer }

func (l loggingCompleter) Complete(ctx context.Context, prompt string) (string, error) {
	start := time.Now()
	out, err := l.next.Complete(ctx, prompt)
	log.Printf("[llm %v] respuesta=%q err=%v", time.Since(start).Round(time.Millisecond), out, err)
	return out, err
}

type loggingToolCaller struct{ next ToolCaller }

func (l loggingToolCaller) CallTool(ctx context.Context, system, user string, tool ToolSpec) (string, error) {
	start := time.Now()
	out, err := l.next.CallTool(ctx, system, user, tool)
	log.Printf("[llm tool %v] args=%q err=%v", time.Since(start).Round(time.Millisecond), out, err)
	return out, err
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

func saveReservation(path string) Listener {
	return func(_ context.Context, e Event) {
		if e.Type != FormCompleted {
			return
		}
		line, err := json.Marshal(map[string]any{
			"session": e.SessionID,
			"at":      time.Now().Format(time.RFC3339),
			"answers": e.Answers,
		})
		if err != nil {
			log.Println("serializando reserva:", err)
			return
		}
		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			log.Println("abriendo archivo de reservas:", err)
			return
		}
		defer f.Close()
		fmt.Fprintln(f, string(line))
	}
}

func NewBotLLM() *Processor {
	today := time.Now()

	// Infraestructura: cliente del LLM y estrategia de extracción (+ logging opcional).
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

	// Dominio: el formulario es solo datos.
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

	// Composición: estrategia LLM + primer mensaje + confirmación + observer.
	p := NewProcessor(form,
		WithFormExtractor(extractor),
		ProcessFirstMessage(),
		RequireConfirmation(),
		OnEvent(saveReservation(getenv("OUT", "reservas.jsonl"))),
	)
	return p
}
