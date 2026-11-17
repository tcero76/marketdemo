package sales

import (
	"context"

	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	dto "github.com/tcero76/marketplace/bff-service/dto/sales"
	"github.com/tcero76/marketplace/bff-service/services"
	inventoryModel "github.com/tcero76/marketplace/postgres/model/inventory"
)

var _ services.IInventory = (*InventoryService)(nil)

type InventoryService struct {
	db *gorm.DB
}

func NewInventoryService(db *gorm.DB) *InventoryService {
	return &InventoryService{
		db: db,
	}
}
func (s *InventoryService) GetStock(
	ctx context.Context,
	productID uuid.UUID,
) (dto.StockResponse, error) {

	return dto.StockResponse{}, nil
}

func (s *InventoryService) Purchase(
	ctx context.Context,
	request dto.PurchaseRequest,
) error {
	return nil
}

func (s *InventoryService) Sell(
	ctx context.Context,
	request dto.SellRequest,
) error {

	return nil
}

func (s *InventoryService) AdjustStock(
	ctx context.Context,
	request dto.AdjustStockRequest,
) error {
	return nil
}

func intAbs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}

func (s *InventoryService) DisableProduct(
	ctx context.Context,
	productID uuid.UUID,
) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var product inventoryModel.Product

		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&product, "id = ?", productID).
			Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrProductNotFound
		}

		if err != nil {
			return err
		}

		if !product.Active {
			return nil
		}

		product.Active = false

		return tx.Save(&product).Error
	})
}

func (s *InventoryService) EnableProduct(
	ctx context.Context,
	productID uuid.UUID,
) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var product inventoryModel.Product

		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&product, "id = ?", productID).
			Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrProductNotFound
		}

		if err != nil {
			return err
		}

		if product.Active {
			return nil
		}

		product.Active = true

		return tx.Save(&product).Error
	})
}
