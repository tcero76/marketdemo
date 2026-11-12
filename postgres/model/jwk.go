package model

import "time"

type JWK struct {
	Kid            string    `gorm:"primaryKey;column:kid"`
	Kty            string    `gorm:"column:kty;not null"`
	Alg            string    `gorm:"column:alg;not null"`
	Use            string    `gorm:"column:use;not null"`
	PrivateJWKJSON []byte    `gorm:"column:private_jwk;type:jsonb;not null"`
	PublicJWKJSON  []byte    `gorm:"column:public_jwk;type:jsonb;not null"`
	CreatedAt      time.Time `gorm:"column:created_at;not null"`
	ExpiresAt      time.Time `gorm:"column:expires_at"`
}

func (JWK) TableName() string {
	return "marketplace.jwk_keys"
}
