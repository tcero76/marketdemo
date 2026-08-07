package oauth2

import (
	"os"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
)

type TokenValidator interface {
	Validate(token string) (jwt.MapClaims, error)
}

type HydraTokenValidator struct {
	jwks *keyfunc.JWKS
}

func NewHydraTokenValidator() *HydraTokenValidator {
	jwksURL := os.Getenv("HYDRA_PUBLIC_URL") + "/.well-known/jwks.json"

	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{})
	if err != nil {
		panic(err)
	}

	return &HydraTokenValidator{
		jwks: jwks,
	}
}

func (v *HydraTokenValidator) Validate(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, v.jwks.Keyfunc)
	if err != nil || !token.Valid {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)

	return claims, nil
}
