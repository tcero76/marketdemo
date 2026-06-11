package dto

import (
	"github.com/google/uuid"
)

type UserDTO struct {
	UserID     uuid.UUID `json:"user_id"`
	Password   string    `json:"password"`
	Provider   string    `json:"provider"`
	Email      string    `json:"email"`
	Avatar     string    `json:"avatar"`
	Nombre     string    `json:"nombre"`
	IDProvider string    `json:"id_provider"`
	Roles      string    `json:"roles"`
}
