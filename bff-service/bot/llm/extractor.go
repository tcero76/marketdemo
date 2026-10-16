package llm

import (
	"context"
	"errors"
	"regexp"
	"strings"
)

// ErrNotUnderstood indica un error recuperable: se vuelve a preguntar al usuario.
var ErrNotUnderstood = errors.New("No pude entender ese dato")

var EmailRe = regexp.MustCompile(`[\w.+-]+@[\w-]+\.[\w.-]+`)

// Extractor (Strategy) convierte un mensaje libre en el valor de un campo.
// Los errores de red/sistema deben devolverse tal cual; ErrNotUnderstood
// es el único que significa "pregúntale de nuevo al usuario".
type Extractor interface {
	Extract(ctx context.Context, input string) (string, error)
}

type ExtractorFunc func(ctx context.Context, input string) (string, error)

func (f ExtractorFunc) Extract(ctx context.Context, in string) (string, error) { return f(ctx, in) }

// TextExtractor toma el mensaje completo como valor.
type TextExtractor struct{}

func (TextExtractor) Extract(_ context.Context, in string) (string, error) {
	v := strings.TrimSpace(in)
	if v == "" {
		return "", ErrNotUnderstood
	}
	return v, nil
}

// RegexExtractor busca la primera coincidencia dentro del mensaje.
type RegexExtractor struct{ Re *regexp.Regexp }

func (e RegexExtractor) Extract(_ context.Context, in string) (string, error) {
	if m := e.Re.FindString(in); m != "" {
		return m, nil
	}
	return "", ErrNotUnderstood
}

// ChoiceExtractor acepta una de varias opciones cerradas.
type ChoiceExtractor struct{ Options []string }

func (e ChoiceExtractor) Extract(_ context.Context, in string) (string, error) {
	low := strings.ToLower(in)
	for _, o := range e.Options {
		if strings.Contains(low, strings.ToLower(o)) {
			return o, nil
		}
	}
	return "", ErrNotUnderstood
}

// Completer es el puerto hacia tu LLM; tú lo implementas con el SDK que uses.
type Completer interface {
	Complete(ctx context.Context, prompt string) (string, error)
}
