package dto

import (
	"github.com/google/uuid"
)

type PostRequest struct {
	ID        uuid.UUID     `json:"id,omitempty"`
	Texto     string        `json:"texto,omitempty"`
	ReplyToID *uuid.UUID    `json:"replyToId,omitempty"`
	ReplyTo   *PostRequest  `json:"replyTo,omitempty"`
	Replies   []PostRequest `json:"replies,omitempty"`
	ProductId uuid.UUID     `json:"productId,omitempty"`
}

type PostResponse struct {
	ID        uuid.UUID      `json:"id,omitempty"`
	Texto     string         `json:"texto,omitempty"`
	ReplyToID *uuid.UUID     `json:"replyToId,omitempty"`
	ReplyTo   *PostResponse  `json:"replyTo,omitempty"`
	Replies   []PostResponse `json:"replies,omitempty"`
	ProductId uuid.UUID      `json:"productId,omitempty"`
}
