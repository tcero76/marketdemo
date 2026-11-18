package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type SearchRequest struct {
	Mention string   `json:"mention"`
	Hashtag string   `json:"hashtag"`
	Text    []string `json:"text"`
}

type ProductResponse struct {
	ID                 uuid.UUID       `json:"id",omitempty`
	Title              string          `json:"title",omitempty`
	Description        string          `json:"description",omitempty`
	Price              decimal.Decimal `json:"price",omitempty`
	DiscountPercentage float64         `json:"discountPercentage",omitempty`
	Rating             float64         `json:"rating",omitempty`
	Quantity           int             `json:"quantity",omitempty`
	Brand              string          `json:"brand",omitempty`
	Category           string          `json:"category",omitempty`
	Thumbnail          string          `json:"thumbnail",omitempty`
}
