package services

import (
	"context"

	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/tcero76/marketplace/bff-service/dto"
	"github.com/tcero76/marketplace/bff-service/services"
	logger "github.com/tcero76/marketplace/config/log"
	"github.com/tcero76/marketplace/postgres/adapters"
	inventoryModel "github.com/tcero76/marketplace/postgres/model/inventory"
)

var _ services.IInventory = (*InventoryService)(nil)

type InventoryService struct {
	dbRead  *gorm.DB
	dbWrite *gorm.DB
	log     *logger.LoggerLogstash
}

func NewInventoryService(log *logger.LoggerLogstash, dbRead *gorm.DB, dbWrite *gorm.DB) *InventoryService {
	return &InventoryService{
		dbRead:  dbRead,
		dbWrite: dbWrite,
		log:     log,
	}
}
func (s *InventoryService) GetStock(ctx context.Context, productID uuid.UUID) (dto.StockResponse, error) {
	return dto.StockResponse{}, nil
}

func (s *InventoryService) Purchase(ctx context.Context, request dto.PurchaseRequest) error {
	return nil
}

func (s *InventoryService) Sell(ctx context.Context, request dto.SellRequest) error {
	return nil
}

func (s *InventoryService) AdjustStock(ctx context.Context, request dto.AdjustStockRequest) error {
	return nil
}

func intAbs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}

func (s *InventoryService) DisableProduct(ctx context.Context, productID uuid.UUID) error {
	return s.dbWrite.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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

func (s *InventoryService) EnableProduct(ctx context.Context, productID uuid.UUID) error {
	return s.dbWrite.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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

func (h *InventoryService) GetCategories() ([]string, error) {
	h.log.Info("Fetching product categories")
	var categories []string
	err := h.dbRead.
		Model(&inventoryModel.Category{}).
		Pluck("name", &categories).Error
	if err != nil {
		h.log.Error("Error fetching categories: ", err)
		return nil, err
	}
	h.log.Debug("Categories found: ", categories)
	return categories, nil
}

func getCategoryIdByName(name string, db *gorm.DB) (uuid.UUID, error) {
	var category inventoryModel.Category
	err := db.
		Where("name = ?", name).
		First(&category).Error
	if err != nil {
		return uuid.Nil, err
	}
	return category.ID, nil
}

func (h *InventoryService) GetProduct(query string) (*dto.ProductResponse, error) {
	h.log.Infof("Fetching product with query: %s", query)
	var product inventoryModel.Product
	err := h.dbRead.
		Preload("Category").Where("id = ?", query).First(&product).Error
	if err != nil {
		h.log.Error("Error fetching product: ", err)
		return nil, err
	}
	return adapters.ToProductDTO(&product), nil
}

func (h *InventoryService) GetProducts() ([]dto.ProductResponse, error) {
	h.log.Infof("Fetching all products")
	var products []inventoryModel.Product
	err := h.dbRead.Find(&products).Error
	if err != nil {
		h.log.Error("Error fetching products: ", err)
		return nil, err
	}
	return adapters.ToProductDTOs(products), nil
}

func (h *InventoryService) GetSearchProduct(searchRequest dto.SearchRequest) []dto.ProductResponse {
	h.log.Debug("Searching products with request: ", searchRequest)
	var productIDs []int
	var categoryIDs []uuid.UUID
	if searchRequest.Hashtag != "" {
		h.log.Debug("Searching by hashtag: %s", searchRequest.Hashtag)
		idCategory, err := getCategoryIdByName(searchRequest.Hashtag, h.dbRead)
		categoryIDs = []uuid.UUID{idCategory}
		if err != nil {
			h.log.Error("No encontró la categoria", err.Error())
		}
	} else {
		h.log.Debug("No hashtag provided, returning empty search results")
	}
	h.log.Debug("Search product IDs: %v", productIDs)
	h.log.Debug("Search category IDs: %v", categoryIDs)
	specs := []Specification{
		TextSpecProduct{Words: searchRequest.Text, Log: h.log},
		CategorySpecProduct{CategoryIDs: categoryIDs, Log: h.log},
	}
	var products []inventoryModel.Product
	err := ApplySpecifications(h.dbRead.Model(&inventoryModel.Product{}), specs...).Find(&products).Error
	if err != nil {
		h.log.Error("Error executing product search: %v", err)
		return nil
	}
	h.log.Debug("Found products: %v", products)
	return adapters.ToProductDTOs(products)
}

func (h *InventoryService) GetRecomendationsProduct(ctx context.Context, userId string) []dto.ProductResponse {
	h.log.Info("Fetching product recommendations for user IDs: ", userId)
	var products []inventoryModel.Product
	err := h.dbRead.
		WithContext(ctx).
		Table("inventory.products").
		Find(&products).Error
	if err != nil {
		h.log.Error("Error fetching products for recommendations: ", err)
		return nil
	}
	h.log.Debug("Products found for recommendations: ", products)
	return adapters.ToProductDTOs(products)
}
