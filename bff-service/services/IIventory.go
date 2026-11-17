package services

import (
	"context"

	"github.com/google/uuid"
	dto "github.com/tcero76/marketplace/bff-service/dto/sales"
)

type IInventory interface {
	GetStock(ctx context.Context, productID uuid.UUID) (dto.StockResponse, error)
	Purchase(ctx context.Context, request dto.PurchaseRequest) error
	Sell(ctx context.Context, request dto.SellRequest) error
	AdjustStock(ctx context.Context, request dto.AdjustStockRequest) error
	DisableProduct(ctx context.Context, productID uuid.UUID) error
	EnableProduct(ctx context.Context, productID uuid.UUID) error
}
