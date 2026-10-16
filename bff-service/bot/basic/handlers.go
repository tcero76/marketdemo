package basic

import (
	"context"
	"regexp"
)

type Input struct {
	Text    string
	History []Message
}

type Reply struct {
	Text string
}

// Todo "cerebro" implementa esta interfaz: reglas hoy, LLM mañana.
type Handler interface {
	Handle(ctx context.Context, in Input) (Reply, error)
}

var (
	reGreeting = regexp.MustCompile(`(?i)\b(hola|buenas)\b`)
	reSchedule = regexp.MustCompile(`(?i)horario`)
)

type RuleBased struct{}

func (RuleBased) Handle(_ context.Context, in Input) (Reply, error) {
	switch {
	case reGreeting.MatchString(in.Text):
		return Reply{Text: "¡Hola! ¿En qué puedo ayudarte?"}, nil
	case reSchedule.MatchString(in.Text):
		return Reply{Text: "Atendemos de lunes a viernes, 9 a 18 h."}, nil
	default:
		return Reply{Text: "No te entendí bien. ¿Puedes reformularlo?"}, nil
	}
}

// Mañana: type LLMHandler struct{ ... } con el mismo método Handle.
