package sales

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreatePurchaseInput struct {
	SupplierID uuid.UUID `json:"supplier_id"`
	UserID     uuid.UUID `json:"user_id"`
	Currency   string    `json:"currency"`
}

type AddPurchaseProductInput struct {
	ProductID uuid.UUID       `json:"product_id"`
	Quantity  int32           `json:"quantity"`
	UnitPrice decimal.Decimal `json:"unit_price"`
}

type CreateReturnPurchaseInput struct {
	PurchaseID uuid.UUID                  `json:"purchase_id"`
	UserID     uuid.UUID                  `json:"user_id"`
	Items      []CreateReturnPurchaseItem `json:"items"`
}

type CreateReturnPurchaseItem struct {
	PurchaseItemID uuid.UUID `json:"purchase_item_id"`
	Quantity       int32     `json:"quantity"`
}
