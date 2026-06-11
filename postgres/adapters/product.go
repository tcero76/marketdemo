package adapters

import (
	"github.com/tcero76/marketplace/bff-service/dto/demo"
	"github.com/tcero76/marketplace/postgres/model"
)

func ToProductDTO(product *model.Product) *demo.Product {
	var brand string
	if product.Brand != nil {
		brand = *product.Brand
	}
	var category string
	if product.Category != nil {
		category = product.Category.Name
	}
	return &demo.Product{
		ID:                 product.ID,
		Title:              product.Title,
		Description:        product.Description,
		Price:              product.Price,
		DiscountPercentage: product.DiscountPercentage,
		Rating:             product.Rating,
		Stock:              product.Stock,
		Brand:              brand,
		Category:           category,
		Thumbnail:          product.Thumbnail,
	}
}

func ToProductDTOs(products []model.Product) []demo.Product {
	var dtos []demo.Product
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
