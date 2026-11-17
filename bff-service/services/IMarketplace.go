package services

import (
	"github.com/tcero76/marketplace/bff-service/dto"
)

type IUserService interface {
	GetUser(username string) (*dto.UserDTO, error)
	GetUserById(userId string) (*dto.UserDTO, error)
	CreateUser(userDTO *dto.UserDTO) error
	GetUserByEmail(email string) (*dto.UserDTO, error)
}
