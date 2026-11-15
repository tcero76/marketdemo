package sales

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/tcero76/marketplace/postgres/model"
	"gorm.io/gorm"
)

type Product struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SKU                string    `gorm:"type:varchar(100);not null;uniqueIndex"`
	Title              string
	Description        string
	Price              decimal.Decimal `gorm:"type:numeric"`
	DiscountPercentage float64         `gorm:"column:discount_percentage;type:numeric"`
	Rating             float64         `gorm:"type:numeric"`
	Stock              int
	Brand              *string
	CategoryID         *uuid.UUID      `gorm:"column:id"`
	Category           *model.Category `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Thumbnail          string
	SearchVector       string               `gorm:"type:tsvector;->"`
	DeletedAt          gorm.DeletedAt       `gorm:"column:deleted_at;index"`
	Images             []model.ProductImage `gorm:"foreignKey:ProductID"`
	Active             bool                 `gorm:"not null;default:true"`
	CreatedAt          time.Time            `gorm:"not null;default:now()"`
}

func (Product) TableName() string {
	return "inventory.products"
}
