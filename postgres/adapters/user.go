package adapters

import (
	"fmt"

	"github.com/tcero76/marketplace/bff-service/dto"
	model "github.com/tcero76/marketplace/postgres/model/marketplace"
)

func ToUserDTO(user model.User) dto.UserResponse {
	var provider string
	switch user.Provider {
	case model.ProviderGoogle:
		provider = "google"
	case model.ProviderInternal:
		provider = "internal"
	default:
		fmt.Println("provider inválido: %s", user.Provider)
	}
	return dto.UserResponse{
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

func ToUser(userRequest dto.UserRequest) model.User {
	var provider model.Provider
	switch userRequest.Provider {
	case "google":
		provider = model.ProviderGoogle
	case "internal":
		provider = model.ProviderInternal
	default:
		fmt.Println("provider inválido: %s", userRequest.Provider)
	}
	return model.User{
		UserID:     userRequest.UserID,
		Nombre:     userRequest.Nombre,
		Email:      userRequest.Email,
		Password:   userRequest.Password,
		Roles:      userRequest.Roles,
		Avatar:     userRequest.Avatar,
		Provider:   provider,
		IDProvider: userRequest.IDProvider,
	}
}
