package llm

import (
	"context"
	"errors"
	"sync"
)

var ErrSessionNotFound = errors.New("sesión no encontrada")

// Phase (State) indica en qué fase está el diálogo. El valor cero es "recolectando".
type Phase string

const (
	PhaseCollecting Phase = ""
	PhaseConfirming Phase = "confirming" // mostró el resumen y espera sí/no
	PhaseCorrecting Phase = "correcting" // el usuario dijo "no" y debe elegir qué dato corregir
)

// Session es el estado serializable de una conversación.
type Session struct {
	ID      string
	FormID  string
	Answers Answers
	History []string // orden en que se llenaron los campos (para "volver")
	Done    bool
	Phase   Phase // fase del diálogo
}

func NewSession(id, formID string) *Session {
	return &Session{ID: id, FormID: formID, Answers: Answers{}}
}

func (s *Session) clone() *Session {
	c := *s
	c.Answers = make(Answers, len(s.Answers))
	for k, v := range s.Answers {
		c.Answers[k] = v
	}
	c.History = append([]string(nil), s.History...)
	return &c
}

// Store (Repository) abstrae dónde viven las sesiones. Cambiar a Redis o
// Postgres = escribir otra implementación; el motor no se entera.
type Store interface {
	Load(ctx context.Context, id string) (*Session, error) // ErrSessionNotFound si no existe
	Save(ctx context.Context, s *Session) error
}

type MemoryStore struct {
	mu   sync.RWMutex
	data map[string]*Session
}

func NewMemoryStore() *MemoryStore { return &MemoryStore{data: map[string]*Session{}} }

func (m *MemoryStore) Load(_ context.Context, id string) (*Session, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	s, ok := m.data[id]
	if !ok {
		return nil, ErrSessionNotFound
	}
	return s.clone(), nil
}

func (m *MemoryStore) Save(_ context.Context, s *Session) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[s.ID] = s.clone()
	return nil
}
