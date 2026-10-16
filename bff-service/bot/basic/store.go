package basic

import (
	"context"
	"sync"
)

type Message struct {
	Role string `json:"role"` // "user" | "bot"
	Text string `json:"text"`
}

type Store interface {
	History(ctx context.Context, convID string) ([]Message, error)
	Append(ctx context.Context, convID string, m Message) error
}

// MemoryStore: válido para empezar. Luego lo cambias por Redis/Postgres.
type MemoryStore struct {
	mu   sync.RWMutex
	data map[string][]Message
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[string][]Message)}
}

func (s *MemoryStore) History(_ context.Context, convID string) ([]Message, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	h := s.data[convID]
	out := make([]Message, len(h)) // copia, para no exponer el slice interno
	copy(out, h)
	return out, nil
}

func (s *MemoryStore) Append(_ context.Context, convID string, m Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[convID] = append(s.data[convID], m)
	return nil
}
