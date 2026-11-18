package services

import (
	"context"

	"github.com/tcero76/marketplace/bff-service/dto"
)

type IPayment interface {
	GetPayment(ctx context.Context, req dto.GetPaymentRequest) (dto.GetPaymentResponse, error)
	Pay(ctx context.Context, req dto.PayRequest) (dto.PayResponse, error)
	Refund(ctx context.Context, req dto.RefundRequest) (dto.RefundResponse, error)
}
