package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID       uint      `gorm:"primaryKey;autoIncrement"`
	Name     string    `gorm:"uniqueIndex;not null"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}

func (Category) TableName() string {
	return "marketplacedemo.category"
}

type Product struct {
	ID                 int `gorm:"primaryKey"`
	Title              string
	Description        string
	Price              float64 `gorm:"type:numeric"`
	DiscountPercentage float64 `gorm:"column:discount_percentage;type:numeric"`
	Rating             float64 `gorm:"type:numeric"`
	Stock              int
	Brand              *string
	CategoryID         *uint     `gorm:"column:categoryid"`
	Category           *Category `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Thumbnail          string
	SearchVector       string         `gorm:"type:tsvector;->"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at;index"`
	Images             []ProductImage `gorm:"foreignKey:ProductID"`
}

func (Product) TableName() string {
	return "marketplacedemo.products"
}

type ProductImage struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	ProductID int  `gorm:"index"`
	URL       string

	Product Product `gorm:"constraint:OnDelete:CASCADE;"`
}

func (ProductImage) TableName() string {
	return "marketplacedemo.product_images"
}

type PostsDemo struct {
	ID         uuid.UUID   `gorm:"type:uuid;primaryKey"`
	Content    string      `gorm:"type:text;not null"`
	ContentTSV string      `gorm:"type:tsvector;->"`
	CreatedAt  time.Time   `gorm:"type:timestamptz;not null;default:now()"`
	UpdatedAt  time.Time   `gorm:"type:timestamptz;not null;default:now()"`
	ReplyToID  *uuid.UUID  `gorm:"type:uuid"`
	ReplyTo    *PostsDemo  `gorm:"foreignKey:ReplyToID;constraint:OnDelete:SET NULL"`
	Replies    []PostsDemo `gorm:"foreignKey:ReplyToID"`
	ProductId  int         `gorm:"index"`
	UserID     uuid.UUID   `gorm:"type:uuid;index"`
}

func (PostsDemo) TableName() string {
	return "marketplacedemo.posts"
}

// func (p *PostsDemo) BeforeCreate(tx *gorm.DB) error {
// 	if p.ID == uuid.Nil {
// 		p.ID = uuid.New()
// 	}
// 	return nil
// }
