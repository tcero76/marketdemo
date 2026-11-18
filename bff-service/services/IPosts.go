package services

import "github.com/tcero76/marketplace/bff-service/dto"

type IDemoPosts interface {
	GetPosts(limit int, offset int) []dto.PostResponse
	GetTotalPosts() int64
	CreatePosteo(posteo *dto.PostRequest, userId string) error
	GetPosteos(productId string) ([]dto.PostResponse, error)
}
