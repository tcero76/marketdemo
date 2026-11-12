package services

import (
	"context"
	"crypto/rsa"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/tcero76/marketplace/bff-service/dto"
	"github.com/tcero76/marketplace/bff-service/dto/demo"
	"github.com/tcero76/marketplace/bff-service/payload"
)

type IUserService interface {
	GetUser(username string) (*dto.UserDTO, error)
	GetUserById(userId string) (*dto.UserDTO, error)
	CreateUser(userDTO *dto.UserDTO) error
	GetUserByEmail(email string) (*dto.UserDTO, error)
}

type IJWKService interface {
	GetKeys() (*rsa.PrivateKey, *jwk.Set, error)
}

type IDemoService interface {
	GetProduct(query string) (*demo.Product, error)
	GetProducts() ([]demo.Product, error)
	GetSearchProduct(searchRequest payload.SearchRequest) []demo.Product
	GetRecomendationsProduct(ctx context.Context, userId string) []demo.Product
	GetCategories() ([]string, error)
}

type IDemoPosts interface {
	GetPosts(limit int, offset int) []demo.PostDTO
	GetTotalPosts() int64
	CreatePosteo(posteo *demo.PostDTO, userId string) error
	GetPosteos(productId string) ([]demo.PostDTO, error)
}
