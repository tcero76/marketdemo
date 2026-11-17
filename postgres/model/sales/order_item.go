package sales

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	inventoryModel "github.com/tcero76/marketplace/postgres/model/inventory"
	schedulingModel "github.com/tcero76/marketplace/postgres/model/scheduling"
)

type OrderItem struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OrderID       uuid.UUID  `gorm:"type:uuid;not null"`
	AppointmentID *uuid.UUID `gorm:"type:uuid"`
	ProductID     *uuid.UUID `gorm:"type:uuid"`

	Quantity    int             `gorm:"not null;default:1"`
	UnitPrice   decimal.Decimal `gorm:"type:numeric(12,2);not null"`
	TotalAmount decimal.Decimal `gorm:"type:numeric(12,2);not null"`
	CreatedAt   time.Time       `gorm:"not null;default:now()"`

	Order       Order                        `gorm:"foreignKey:OrderID;references:ID"`
	Appointment *schedulingModel.Appointment `gorm:"foreignKey:AppointmentID;references:ID"`
	Product     *inventoryModel.Product      `gorm:"foreignKey:ProductID;references:ID"`
}

func (OrderItem) TableName() string {
	return "sales.order_items"
}
