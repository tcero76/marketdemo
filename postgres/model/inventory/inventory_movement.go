package inventory

import (
	"time"

	"github.com/google/uuid"
	model "github.com/tcero76/marketplace/postgres/model/marketplace"
)

type MovementType string

const (
	MovementPurchase   MovementType = "PURCHASE"
	MovementSale       MovementType = "SALE"
	MovementAdjustment MovementType = "ADJUSTMENT"
)

type InventoryMovement struct {
	ID           uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProductID    uuid.UUID    `gorm:"type:uuid;not null"`
	UserID       uuid.UUID    `gorm:"type:uuid;not null"`
	OrderID      *uuid.UUID   `gorm:"type:uuid;not null"`
	MovementType MovementType `gorm:"type:varchar(20);not null"`
	Quantity     int          `gorm:"not null"`
	CreatedAt    time.Time    `gorm:"not null;default:now()"`

	Product Product    `gorm:"foreignKey:ProductID;references:ID"`
	User    model.User `gorm:"foreignKey:UserID;references:ID"`
}

func (InventoryMovement) TableName() string {
	return "inventory.inventory_movements"
}
