package llm

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

// ValidationError es un error recuperable que se muestra al usuario.
type ValidationError struct{ Msg string }

func (e *ValidationError) Error() string { return e.Msg }

func Invalid(format string, args ...any) error {
	return &ValidationError{Msg: fmt.Sprintf(format, args...)}
}

// Validator (Strategy) verifica un valor ya extraído.
type Validator interface {
	Validate(value string) error
}

type ValidatorFunc func(string) error

func (f ValidatorFunc) Validate(v string) error { return f(v) }

func MinLength(n int) Validator {
	return ValidatorFunc(func(v string) error {
		if utf8.RuneCountInString(v) < n {
			return Invalid("Debe tener al menos %d caracteres", n)
		}
		return nil
	})
}

func Matches(re *regexp.Regexp, msg string) Validator {
	return ValidatorFunc(func(v string) error {
		if !re.MatchString(v) {
			return Invalid("%s", msg)
		}
		return nil
	})
}

func OneOf(opts ...string) Validator {
	return ValidatorFunc(func(v string) error {
		for _, o := range opts {
			if v == o {
				return nil
			}
		}
		return Invalid("Debe ser una de estas opciones: %s", strings.Join(opts, ", "))
	})
}

func runValidators(vs []Validator, value string) error {
	for _, v := range vs {
		if err := v.Validate(value); err != nil {
			return err
		}
	}
	return nil
}
