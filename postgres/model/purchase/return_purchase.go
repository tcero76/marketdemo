package purchase

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ReturnPurchaseStatus string

const (
	ReturnPurchaseStatusPending   ReturnPurchaseStatus = "PENDING"
	ReturnPurchaseStatusCancelled ReturnPurchaseStatus = "CANCELLED"
	ReturnPurchaseStatusCompleted ReturnPurchaseStatus = "COMPLETED"
)

type ReturnPurchase struct {
	ID          uuid.UUID            `json:"id"`
	PurchaseID  uuid.UUID            `json:"purchase_id"`
	UserID      uuid.UUID            `json:"user_id"`
	Status      ReturnPurchaseStatus `json:"status"`
	TotalAmount decimal.Decimal      `json:"total_amount"`
	Currency    string               `json:"currency"`
	Items       []ReturnPurchaseItem `json:"items"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

func (posts ReturnPurchase) TableName() string {
	return "purchasing.return_purchases"
}
