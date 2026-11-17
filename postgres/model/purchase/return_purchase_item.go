package purchase

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ReturnPurchaseItem struct {
	ID               uuid.UUID       `json:"id"`
	ReturnPurchaseID uuid.UUID       `json:"return_purchase_id"`
	PurchaseItemID   uuid.UUID       `json:"purchase_item_id"`
	Quantity         int32           `json:"quantity"`
	UnitPrice        decimal.Decimal `json:"unit_price"`
	TotalAmount      decimal.Decimal `json:"total_amount"`
	CreatedAt        time.Time       `json:"created_at"`
}

func (posts ReturnPurchaseItem) TableName() string {
	return "purchasing.return_purchase_items"
}
