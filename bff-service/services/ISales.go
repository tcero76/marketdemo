package services

import (
	"context"

	dto "github.com/tcero76/marketplace/bff-service/dto"
	"github.com/tcero76/marketplace/bff-service/dto/demo"
	salesDto "github.com/tcero76/marketplace/bff-service/dto/sales"
)

type IDemoService interface {
	GetProduct(query string) (*demo.Product, error)
	GetProducts() ([]demo.Product, error)
	GetSearchProduct(searchRequest dto.SearchRequest) []demo.Product
	GetRecomendationsProduct(ctx context.Context, userId string) []demo.Product
	GetCategories() ([]string, error)
}

type ISales interface {
	GetOrder(ctx context.Context, req salesDto.GetOrderRequest) (salesDto.GetOrderResponse, error)
	CreateOrder(ctx context.Context, req salesDto.CreateOrderRequest) (salesDto.CreateOrderResponse, error)
	AddProduct(ctx context.Context, req salesDto.AddProductRequest) (salesDto.AddProductResponse, error)
	AddAppointment(ctx context.Context, req salesDto.AddAppointmentRequest) (salesDto.AddAppointmentResponse, error)
	// PayOrder(ctx context.Context, req salesDto.PayOrderRequest) (salesDto.PayOrderResponse, error)
	CancelOrder(ctx context.Context, req salesDto.CancelOrderRequest) (salesDto.CancelOrderResponse, error)
	CreateReturn(ctx context.Context, req salesDto.CreateReturnRequest) (salesDto.CreateReturnResponse, error)
}
