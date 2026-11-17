package demo

import (
	"context"

	"github.com/google/uuid"
	bffDemo "github.com/tcero76/marketplace/bff-service/dto/demo"
	"github.com/tcero76/marketplace/bff-service/payload"
	"github.com/tcero76/marketplace/bff-service/services"
	logger "github.com/tcero76/marketplace/config/log"
	"github.com/tcero76/marketplace/postgres/adapters"
	inventoryModel "github.com/tcero76/marketplace/postgres/model/inventory"
	"gorm.io/gorm"
)

type ProductService struct {
	dbWrite *gorm.DB
	dbRead  *gorm.DB
	log     *logger.LoggerLogstash
}

func NewProductService(log *logger.LoggerLogstash, dbWrite *gorm.DB, dbRead *gorm.DB) services.IDemoService {
	return &ProductService{dbWrite: dbWrite, dbRead: dbRead, log: log}
}

func (h *ProductService) GetProduct(query string) (*bffDemo.Product, error) {
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

func (h *ProductService) GetProducts() ([]bffDemo.Product, error) {
	h.log.Infof("Fetching all products")
	var products []inventoryModel.Product
	err := h.dbRead.Find(&products).Error
	if err != nil {
		h.log.Error("Error fetching products: ", err)
		return nil, err
	}
	return adapters.ToProductDTOs(products), nil
}

func (h *ProductService) GetSearchProduct(searchRequest payload.SearchRequest) []bffDemo.Product {
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

func (h *ProductService) GetRecomendationsProduct(ctx context.Context, userId string) []bffDemo.Product {
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

func (h *ProductService) GetCategories() ([]string, error) {
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
