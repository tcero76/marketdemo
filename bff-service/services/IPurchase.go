package services

import (
	"context"

	"github.com/google/uuid"
	salesModel "github.com/tcero76/marketplace/bff-service/dto/sales"
	purchaseModel "github.com/tcero76/marketplace/postgres/model/purchase"
)

type Purchasing interface {
	GetPurchase(ctx context.Context, purchaseID uuid.UUID) (purchaseModel.Purchase, error)
	CreatePurchase(ctx context.Context, input salesModel.CreatePurchaseInput) (purchaseModel.Purchase, error)
	AddProduct(ctx context.Context, purchaseID uuid.UUID, input salesModel.AddPurchaseProductInput) (purchaseModel.PurchaseItem, error)
	ReceivePurchase(ctx context.Context, purchaseID uuid.UUID) error
	CancelPurchase(ctx context.Context, purchaseID uuid.UUID) error
	GetReturnPurchase(ctx context.Context, returnPurchaseID uuid.UUID) (purchaseModel.ReturnPurchase, error)
	CreateReturnPurchase(ctx context.Context, input salesModel.CreateReturnPurchaseInput) (purchaseModel.ReturnPurchase, error)
	CompleteReturnPurchase(ctx context.Context, returnPurchaseID uuid.UUID) error
	CancelReturnPurchase(ctx context.Context, returnPurchaseID uuid.UUID) error
}
