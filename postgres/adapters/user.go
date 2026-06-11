package adapters

import (
	"fmt"

	"github.com/tcero76/marketplace/bff-service/dto"
	"github.com/tcero76/marketplace/postgres/model"
)

func ToUserDTO(user model.User) dto.UserDTO {
	var provider string
	switch user.Provider {
	case model.ProviderGoogle:
		provider = "google"
	case model.ProviderInternal:
		provider = "internal"
	default:
		fmt.Println("provider inválido: %s", user.Provider)
	}
	return dto.UserDTO{
		UserID:     user.UserID,
		Nombre:     user.Nombre,
		Email:      user.Email,
		Password:   user.Password,
		Roles:      user.Roles,
		Avatar:     user.Avatar,
		Provider:   provider,
		IDProvider: user.IDProvider,
	}
}

func ToUser(userDTO dto.UserDTO) model.User {
	var provider model.Provider
	switch userDTO.Provider {
	case "google":
		provider = model.ProviderGoogle
	case "internal":
		provider = model.ProviderInternal
	default:
		fmt.Println("provider inválido: %s", userDTO.Provider)
	}
	return model.User{
		UserID:     userDTO.UserID,
		Nombre:     userDTO.Nombre,
		Email:      userDTO.Email,
		Password:   userDTO.Password,
		Roles:      userDTO.Roles,
		Avatar:     userDTO.Avatar,
		Provider:   provider,
		IDProvider: userDTO.IDProvider,
	}
}
