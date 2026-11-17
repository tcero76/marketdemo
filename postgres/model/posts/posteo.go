package posts

import (
	"time"

	"github.com/google/uuid"
)

// type Posteo struct {
// 	ID        string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
// 	Texto     string         `gorm:"texto"`
// 	Menciones pq.StringArray `gorm:"type:text[]"`
// 	UserId    uuid.UUID      `gorm:"user_id;type:uuid;not null"`
// }

// func (posts Posteo) TableName() string {
// 	return "posts.posteos"
// }

type Posts struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Content    string     `gorm:"type:text;not null"`
	ContentTSV string     `gorm:"type:tsvector;->"`
	CreatedAt  time.Time  `gorm:"type:timestamptz;not null;default:now()"`
	UpdatedAt  time.Time  `gorm:"type:timestamptz;not null;default:now()"`
	ReplyToID  *uuid.UUID `gorm:"type:uuid"`
	ReplyTo    *Posts     `gorm:"foreignKey:ReplyToID;constraint:OnDelete:SET NULL"`
	Replies    []Posts    `gorm:"foreignKey:ReplyToID"`
	ProductId  int        `gorm:"index"`
	UserID     uuid.UUID  `gorm:"type:uuid;index"`
}

func (Posts) TableName() string {
	return "posts.posts"
}
