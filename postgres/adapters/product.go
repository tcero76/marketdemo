package adapters

import (
	"github.com/tcero76/marketplace/bff-service/dto"
	model "github.com/tcero76/marketplace/postgres/model/inventory"
)

func ToProductDTO(product *model.Product) *dto.ProductResponse {
	var brand string
	if product.Brand != nil {
		brand = *product.Brand
	}
	var category string
	if product.Category != nil {
		category = product.Category.Name
	}
	return &dto.ProductResponse{
		ID:                 product.ID,
		Title:              product.Title,
		Description:        product.Description,
		Price:              product.Price,
		DiscountPercentage: product.DiscountPercentage,
		Rating:             product.Rating,
		Quantity:           product.Quantity,
		Brand:              brand,
		Category:           category,
		Thumbnail:          product.Thumbnail,
	}
}

func ToProductDTOs(products []model.Product) []dto.ProductResponse {
	var dtos []dto.ProductResponse
	for _, product := range products {
		dtos = append(dtos, *ToProductDTO(&product))
	}
	return dtos
}

type ProductImages struct {
	ID        int    `json:"id"`
	ProductID int    `json:"productId"`
	URL       string `json:"url"`
}
