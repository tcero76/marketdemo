package basic

import (
	"context"
	"strings"
)

const maxRunes = 2000

type Bot struct {
	store   Store
	handler Handler
}

func New(s Store, h Handler) *Bot {
	return &Bot{store: s, handler: h}
}

func (b *Bot) Process(ctx context.Context, convID, text string) (Reply, error) {
	// 1. normalizar / validar
	text = strings.TrimSpace(text)
	if r := []rune(text); len(r) > maxRunes {
		text = string(r[:maxRunes])
	}
	if text == "" {
		return Reply{Text: "Escribe algo para empezar."}, nil
	}

	// 2. contexto + decidir respuesta
	history, err := b.store.History(ctx, convID)
	if err != nil {
		return Reply{}, err
	}
	if err := b.store.Append(ctx, convID, Message{Role: "user", Text: text}); err != nil {
		return Reply{}, err
	}

	reply, err := b.handler.Handle(ctx, Input{Text: text, History: history})
	if err != nil {
		return Reply{}, err
	}

	// 3. guardar y devolver
	if err := b.store.Append(ctx, convID, Message{Role: "bot", Text: reply.Text}); err != nil {
		return Reply{}, err
	}
	return reply, nil
}