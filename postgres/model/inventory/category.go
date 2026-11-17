package inventory

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string    `gorm:"uniqueIndex;not null"`
}

func (Category) TableName() string {
	return "inventory.category"
}
