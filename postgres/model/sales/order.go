package sales

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	model "github.com/tcero76/marketplace/postgres/model/marketplace"
)

type OrderStatus string

const (
	OrderPending   OrderStatus = "PENDING"
	OrderCancelled OrderStatus = "CANCELLED"
	OrderPaid      OrderStatus = "PAID"
	OrderCompleted OrderStatus = "COMPLETED"
)

type Order struct {
	ID          uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id"`
	UserID      uuid.UUID       `gorm:"type:uuid;not null;column:user_id"`
	Status      OrderStatus     `gorm:"type:varchar(20);not null;default:'PENDING'"`
	TotalAmount decimal.Decimal `gorm:"type:numeric(12,2);not null"`
	Currency    string          `gorm:"type:char(3);not null"`
	CreatedAt   time.Time       `gorm:"not null;default:now()"`
	UpdatedAt   time.Time       `gorm:"not null;default:now()"`

	User  model.User  `gorm:"foreignKey:UserID;references:ID"`
	Items []OrderItem `gorm:"foreignKey:OrderID"`
}

func (Order) TableName() string {
	return "sales.orders"
}
