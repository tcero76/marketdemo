package demo

import (
	"github.com/google/uuid"
)

type PostDTO struct {
	ID        uuid.UUID  `json:"id,omitempty"`
	Texto     string     `json:"texto,omitempty"`
	ReplyToID *uuid.UUID `json:"replyToId,omitempty"`
	ReplyTo   *PostDTO   `json:"replyTo,omitempty"`
	Replies   []PostDTO  `json:"replies,omitempty"`
	ProductId int        `json:"productId,omitempty"`
}
