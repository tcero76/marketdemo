package services

import (
	"context"

	dto "github.com/tcero76/marketplace/bff-service/dto/sales"
)

type IScheduling interface {
	GetSchedule(ctx context.Context, req dto.GetScheduleRequest) (dto.GetScheduleResponse, error)
	GetAppointment(ctx context.Context, req dto.GetAppointmentRequest) (dto.GetAppointmentResponse, error)
	Schedule(ctx context.Context, req dto.ScheduleRequest) (dto.ScheduleResponse, error)
	Reschedule(ctx context.Context, req dto.RescheduleRequest) (dto.RescheduleResponse, error)
	Cancel(ctx context.Context, req dto.CancelRequest) (dto.CancelResponse, error)
}
