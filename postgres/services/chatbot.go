package services

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	model "github.com/tcero76/marketplace/postgres/model/chatbot"
)

const (
	MinPersonas = 1
	MaxPersonas = 10
)

var (
	ErrDatosInvalidos = errors.New("datos de reserva inválidos")
	ErrNoEncontrada   = errors.New("reserva no encontrada")
)

var emailRe = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

type NuevaReserva struct {
	SessionID string
	Nombre    string
	Email     string
	Fecha     time.Time
	Personas  int
}

func FromAnswers(sessionID string, a map[string]string) (NuevaReserva, error) {
	fecha, err := time.Parse("2006-01-02", strings.TrimSpace(a["fecha"]))
	if err != nil {
		return NuevaReserva{}, invalid("fecha %q no tiene formato AAAA-MM-DD", a["fecha"])
	}
	personas, err := strconv.Atoi(strings.TrimSpace(a["personas"]))
	if err != nil {
		return NuevaReserva{}, invalid("personas %q no es un número", a["personas"])
	}
	return NuevaReserva{
		SessionID: sessionID,
		Nombre:    a["nombre"],
		Email:     a["email"],
		Fecha:     fecha,
		Personas:  personas,
	}, nil
}

func invalid(format string, args ...any) error {
	return fmt.Errorf("%w: %s", ErrDatosInvalidos, fmt.Sprintf(format, args...))
}

func (n *NuevaReserva) normalize() error {
	n.SessionID = strings.TrimSpace(n.SessionID)
	n.Nombre = strings.TrimSpace(n.Nombre)
	n.Email = strings.ToLower(strings.TrimSpace(n.Email))

	switch {
	case n.SessionID == "":
		return invalid("session_id vacío")
	case utf8.RuneCountInString(n.Nombre) < 2:
		return invalid("el nombre debe tener al menos 2 caracteres")
	case !emailRe.MatchString(n.Email):
		return invalid("correo %q inválido", n.Email)
	case n.Fecha.IsZero():
		return invalid("falta la fecha")
	case n.Personas < MinPersonas || n.Personas > MaxPersonas:
		return invalid("personas debe estar entre %d y %d", MinPersonas, MaxPersonas)
	}
	y, m, d := n.Fecha.Date()
	n.Fecha = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	return nil
}

type Service struct{ db *gorm.DB }

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}
func (s *Service) Create(ctx context.Context, answers map[string]string, sessionid string) (*model.Reserva, error) {
	fecha, err := time.Parse("2006-01-02", answers["fecha"])
	if err != nil {
		return &model.Reserva{}, err
	}
	personas, err := strconv.ParseInt(answers["personas"], 10, 16)
	if err != nil {
		return &model.Reserva{}, err
	}
	r := model.Reserva{
		SessionID: answers["ID"],
		Nombre:    answers["nombre"],
		Email:     answers["email"],
		Fecha:     fecha,
		Personas:  int16(personas),
	}
	res := s.db.WithContext(ctx).
		Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "session_id"}}, DoNothing: true}).
		Create(&r)
	if res.Error != nil {
		return nil, fmt.Errorf("guardando reserva: %w", res.Error)
	}
	if res.RowsAffected == 0 { // ya existía una reserva para esta conversación
		return s.GetBySession(ctx, sessionid)
	}
	return &r, nil
}

func (s *Service) GetBySession(ctx context.Context, sessionID string) (*model.Reserva, error) {
	var r model.Reserva
	err := s.db.WithContext(ctx).Where("session_id = ?", sessionID).First(&r).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNoEncontrada
	}
	if err != nil {
		return nil, fmt.Errorf("consultando reserva: %w", err)
	}
	return &r, nil
}
