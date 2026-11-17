package services

import (
	"crypto/rsa"

	"github.com/lestrrat-go/jwx/jwk"
)

type IJWKService interface {
	GetKeys() (*rsa.PrivateKey, *jwk.Set, error)
}
