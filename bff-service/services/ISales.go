package services

import (
	"context"

	"github.com/google/uuid"
	dto "github.com/tcero76/marketplace/bff-service/dto/sales"
)

type IInventory interface {
	GetStock(ctx context.Context, productID uuid.UUID) (dto.StockResponse, error)
	Purchase(ctx context.Context, request dto.PurchaseRequest) error
	Sell(ctx context.Context, request dto.SellRequest) error
	AdjustStock(ctx context.Context, request dto.AdjustStockRequest) error
	DisableProduct(ctx context.Context, productID uuid.UUID) error
	EnableProduct(ctx context.Context, productID uuid.UUID) error
}

type IScheduling interface {
	GetSchedule(ctx context.Context, req dto.GetScheduleRequest) (dto.GetScheduleResponse, error)
	GetAppointment(ctx context.Context, req dto.GetAppointmentRequest) (dto.GetAppointmentResponse, error)
	Schedule(ctx context.Context, req dto.ScheduleRequest) (dto.ScheduleResponse, error)
	Reschedule(ctx context.Context, req dto.RescheduleRequest) (dto.RescheduleResponse, error)
	Cancel(ctx context.Context, req dto.CancelRequest) (dto.CancelResponse, error)
}

type IPayment interface {
	GetPayment(ctx context.Context, req dto.GetPaymentRequest) (dto.GetPaymentResponse, error)
	Pay(ctx context.Context, req dto.PayRequest) (dto.PayResponse, error)
	Refund(ctx context.Context, req dto.RefundRequest) (dto.RefundResponse, error)
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
