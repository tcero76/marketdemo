package services

import (
	"github.com/tcero76/marketplace/bff-service/dto"
)

type IUserService interface {
	GetUser(username string) (*dto.UserResponse, error)
	GetUserById(userId string) (*dto.UserResponse, error)
	CreateUser(userDTO *dto.UserRequest) error
	GetUserByEmail(email string) (*dto.UserResponse, error)
}
