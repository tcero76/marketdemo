package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string    `gorm:"uniqueIndex;not null"`
}

func (Category) TableName() string {
	return "inventory.category"
}

type ProductImage struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	ProductID int  `gorm:"index"`
	URL       string
}

func (ProductImage) TableName() string {
	return "inventory.product_images"
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
	return "posts.posts"
}
