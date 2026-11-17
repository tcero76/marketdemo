package purchase

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PurchaseStatus string

const (
	PurchaseStatusPending   PurchaseStatus = "PENDING"
	PurchaseStatusCancelled PurchaseStatus = "CANCELLED"
	PurchaseStatusReceived  PurchaseStatus = "RECEIVED"
)

type Purchase struct {
	ID          uuid.UUID       `json:"id"`
	SupplierID  uuid.UUID       `json:"supplier_id"`
	UserID      uuid.UUID       `json:"user_id"`
	Status      PurchaseStatus  `json:"status"`
	TotalAmount decimal.Decimal `json:"total_amount"`
	Currency    string          `json:"currency"`
	Items       []PurchaseItem  `json:"items"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

func (posts Purchase) TableName() string {
	return "purchasing.purchases"
}
