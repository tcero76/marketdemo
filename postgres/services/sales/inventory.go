package sales

import (
	"context"

	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	dto "github.com/tcero76/marketplace/bff-service/dto/sales"
	"github.com/tcero76/marketplace/bff-service/services"
	"github.com/tcero76/marketplace/postgres/model/sales"
)

var _ services.Inventory = (*InventoryService)(nil)

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
	var stock sales.Stock

	err := s.db.WithContext(ctx).
		Where("product_id = ?", productID).
		First(&stock).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.StockResponse{
				ProductID: productID,
				Quantity:  0,
			}, nil
		}

		return dto.StockResponse{}, err
	}

	return dto.StockResponse{
		ProductID: stock.ProductID,
		Quantity:  stock.Quantity,
	}, nil
}

func (s *InventoryService) Purchase(
	ctx context.Context,
	request dto.PurchaseRequest,
) error {
	if request.Quantity <= 0 {
		return ErrInvalidQuantity
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var product sales.Product
		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&product, "id = ?", request.ProductID).
			Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrProductNotFound
			}

			return err
		}
		if !product.Active {
			return ErrProductDisabled
		}
		var stock sales.Stock
		err = tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&stock, "product_id = ?", request.ProductID).
			Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			stock = sales.Stock{
				ProductID: request.ProductID,
				Quantity:  request.Quantity,
			}

			if err := tx.Create(&stock).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		} else {
			stock.Quantity += request.Quantity

			if err := tx.Save(&stock).Error; err != nil {
				return err
			}
		}

		movement := sales.InventoryMovement{
			ProductID:    request.ProductID,
			UserID:       request.UserID,
			MovementType: sales.MovementPurchase,
			Quantity:     request.Quantity,
		}

		return tx.Create(&movement).Error
	})
}

func (s *InventoryService) Sell(
	ctx context.Context,
	request dto.SellRequest,
) error {
	if request.Quantity <= 0 {
		return ErrInvalidQuantity
	}

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var product sales.Product

		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&product, "id = ?", request.ProductID).
			Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrProductNotFound
			}

			return err
		}

		if !product.Active {
			return ErrProductDisabled
		}

		var stock sales.Stock

		err = tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&stock, "product_id = ?", request.ProductID).
			Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrInsufficientStock
		}

		if err != nil {
			return err
		}

		if stock.Quantity < request.Quantity {
			return ErrInsufficientStock
		}

		stock.Quantity -= request.Quantity

		if err := tx.Save(&stock).Error; err != nil {
			return err
		}

		movement := sales.InventoryMovement{
			ProductID:    request.ProductID,
			UserID:       request.UserID,
			MovementType: sales.MovementSale,
			Quantity:     request.Quantity,
		}

		return tx.Create(&movement).Error
	})
}

func (s *InventoryService) AdjustStock(
	ctx context.Context,
	request dto.AdjustStockRequest,
) error {
	if request.ActualQuantity < 0 {
		return ErrInvalidQuantity
	}

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var product sales.Product

		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&product, "id = ?", request.ProductID).
			Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrProductNotFound
			}

			return err
		}

		if !product.Active {
			return ErrProductDisabled
		}

		var stock sales.Stock

		err = tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&stock, "product_id = ?", request.ProductID).
			Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			stock = sales.Stock{
				ProductID: request.ProductID,
				Quantity:  request.ActualQuantity,
			}

			if err := tx.Create(&stock).Error; err != nil {
				return err
			}

			if request.ActualQuantity == 0 {
				return nil
			}

			return tx.Create(&sales.InventoryMovement{
				ProductID:    request.ProductID,
				UserID:       request.UserID,
				MovementType: sales.MovementAdjustment,
				Quantity:     request.ActualQuantity,
			}).Error
		}

		if err != nil {
			return err
		}

		delta := request.ActualQuantity - stock.Quantity

		if delta == 0 {
			return nil
		}

		stock.Quantity = request.ActualQuantity

		if err := tx.Save(&stock).Error; err != nil {
			return err
		}

		movement := sales.InventoryMovement{
			ProductID:    request.ProductID,
			UserID:       request.UserID,
			MovementType: sales.MovementAdjustment,
			Quantity:     intAbs(delta),
		}

		return tx.Create(&movement).Error
	})
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
		var product sales.Product

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
		var product sales.Product

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
