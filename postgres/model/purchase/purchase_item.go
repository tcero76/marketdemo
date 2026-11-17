package purchase

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PurchaseItem struct {
	ID          uuid.UUID       `json:"id"`
	PurchaseID  uuid.UUID       `json:"purchase_id"`
	ProductID   uuid.UUID       `json:"product_id"`
	Quantity    int32           `json:"quantity"`
	UnitPrice   decimal.Decimal `json:"unit_price"`
	TotalAmount decimal.Decimal `json:"total_amount"`
	CreatedAt   time.Time       `json:"created_at"`
}

func (posts PurchaseItem) TableName() string {
	return "purchasing.purchase_items"
}
