package services

import (
	"context"

	"github.com/tcero76/marketplace/bff-service/dto/demo"
	dto "github.com/tcero76/marketplace/bff-service/dto/sales"
	"github.com/tcero76/marketplace/bff-service/payload"
)

type IDemoService interface {
	GetProduct(query string) (*demo.Product, error)
	GetProducts() ([]demo.Product, error)
	GetSearchProduct(searchRequest payload.SearchRequest) []demo.Product
	GetRecomendationsProduct(ctx context.Context, userId string) []demo.Product
	GetCategories() ([]string, error)
}

type ISales interface {
	GetOrder(ctx context.Context, req dto.GetOrderRequest) (dto.GetOrderResponse, error)
	CreateOrder(ctx context.Context, req dto.CreateOrderRequest) (dto.CreateOrderResponse, error)
	AddProduct(ctx context.Context, req dto.AddProductRequest) (dto.AddProductResponse, error)
	AddAppointment(ctx context.Context, req dto.AddAppointmentRequest) (dto.AddAppointmentResponse, error)
	// PayOrder(ctx context.Context, req dto.PayOrderRequest) (dto.PayOrderResponse, error)
	CancelOrder(ctx context.Context, req dto.CancelOrderRequest) (dto.CancelOrderResponse, error)
	CreateReturn(ctx context.Context, req dto.CreateReturnRequest) (dto.CreateReturnResponse, error)
}
