package services

import "github.com/tcero76/marketplace/bff-service/dto/demo"

type IDemoPosts interface {
	GetPosts(limit int, offset int) []demo.PostDTO
	GetTotalPosts() int64
	CreatePosteo(posteo *demo.PostDTO, userId string) error
	GetPosteos(productId string) ([]demo.PostDTO, error)
}
