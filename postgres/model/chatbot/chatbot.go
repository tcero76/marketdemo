package chatbot

import "time"

type Reserva struct {
	ID        int64     `gorm:"primaryKey"`
	SessionID string    `gorm:"column:session_id;not null"`
	Nombre    string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	Fecha     time.Time `gorm:"type:date;not null"`
	Personas  int16     `gorm:"type:smallint;not null"`
	CreatedAt time.Time `gorm:"not null"`
}

func (Reserva) TableName() string { return "chatbot.reservas" }
