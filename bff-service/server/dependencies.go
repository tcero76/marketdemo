package server

import (
	hydra "github.com/ory/hydra-client-go/v2"
	_ "github.com/tcero76/marketplace/bff-service/docs"
	oauth2Internal "github.com/tcero76/marketplace/bff-service/oauth2"
	"github.com/tcero76/marketplace/bff-service/services"
	"gorm.io/gorm"

	logConfig "github.com/tcero76/marketplace/config/log"
	"golang.org/x/oauth2"
)

type Dependencies struct {
	Log *logConfig.LoggerLogstash

	DBRead  *gorm.DB
	DBWrite *gorm.DB

	AuthCacheService services.IAuthCacheService

	UserService    services.IUserService
	ProductService services.IDemoService
	PostsService   services.IDemoPosts

	JWKService   services.IJWKService
	HydraClient  *hydra.APIClient
	Validator    oauth2Internal.TokenValidator
	GoogleClient *oauth2.Config
}
