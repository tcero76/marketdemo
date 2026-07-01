package demo

import (
	"context"
	"strconv"

	"github.com/tcero76/marketplace/bff-service/dto/demo"
	"github.com/tcero76/marketplace/bff-service/payload"
	logger "github.com/tcero76/marketplace/config/log"
	"github.com/tcero76/marketplace/postgres/adapters"
	"github.com/tcero76/marketplace/postgres/model"
	"gorm.io/gorm"
)

type ProductService struct {
	dbWrite *gorm.DB
	dbRead  *gorm.DB
	log     *logger.LoggerLogstash
}

func NewProductService(log *logger.LoggerLogstash, dbWrite *gorm.DB, dbRead *gorm.DB) *ProductService {
	return &ProductService{dbWrite: dbWrite, dbRead: dbRead, log: log}
}

func (h *ProductService) GetProduct(query string) (*demo.Product, error) {
	h.log.Infof("Fetching product with query: %s", query)
	var product model.Product
	err := h.dbRead.
		Preload("Category").Where("id = ?", query).First(&product).Error
	if err != nil {
		h.log.Error("Error fetching product: ", err)
		return nil, err
	}
	return adapters.ToProductDTO(&product), nil
}
func (h *ProductService) GetProducts() ([]demo.Product, error) {
	h.log.Infof("Fetching all products")
	var products []model.Product
	err := h.dbRead.Find(&products).Error
	if err != nil {
		h.log.Error("Error fetching products: ", err)
		return nil, err
	}
	return adapters.ToProductDTOs(products), nil
}

func (h *ProductService) GetSearchProduct(searchRequest payload.SearchRequest) []demo.Product {
	h.log.Debug("Searching products with request: ", searchRequest)
	var productIDs []int
	if searchRequest.Hashtag != "" {
		h.log.Debug("Searching by hashtag: %s", searchRequest.Hashtag)
		if id, err := strconv.Atoi(searchRequest.Hashtag); err == nil {
			productIDs = []int{id}
		} else {
			h.log.Warn("Invalid hashtag (not a number): %s", searchRequest.Hashtag)
		}
	} else {
		h.log.Debug("No hashtag provided, returning empty search results")
	}
	h.log.Debug("Search product IDs: %v", productIDs)
	specs := []Specification{
		TextSpecProduct{Words: searchRequest.Text, Log: h.log},
	}
	var products []model.Product
	err := ApplySpecifications(h.dbRead.Model(&model.Product{}), specs...).Find(&products).Error
	if err != nil {
		h.log.Error("Error executing product search: %v", err)
		return nil
	}
	h.log.Debug("Found products: %v", products)
	return adapters.ToProductDTOs(products)
}

func (h *ProductService) GetRecomendationsProduct(ctx context.Context, userId string) []int {
	h.log.Info("Fetching product recommendations for user ID: ", userId)
	var products []int
	err := h.dbRead.
		Model(&model.Product{}).
		Select("id").
		Scan(&products).Error
	if err != nil {
		h.log.Error("Error fetching products for recommendations: ", err)
		return nil
	}
	h.log.Debug("Products found for recommendations: ", products)
	return products
}

func (h *ProductService) GetCategories() ([]string, error) {
	h.log.Info("Fetching product categories")
	var categories []string
	err := h.dbRead.
		Model(&model.Category{}).
		Pluck("name", &categories).Error
	if err != nil {
		h.log.Error("Error fetching categories: ", err)
		return nil, err
	}
	h.log.Debug("Categories found: ", categories)
	return categories, nil
}
