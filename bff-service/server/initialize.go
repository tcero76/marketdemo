package server

import (
	"context"
	"os"

	"github.com/golang-jwt/jwt/v4"
	hydra "github.com/ory/hydra-client-go/v2"
	"github.com/tcero76/marketplace/bff-service/config"
	_ "github.com/tcero76/marketplace/bff-service/docs"
	"github.com/tcero76/marketplace/bff-service/oauth2"
	"github.com/tcero76/marketplace/bff-service/services"

	logConfig "github.com/tcero76/marketplace/config/log"

	postgresConfig "github.com/tcero76/marketplace/postgres/config"
	postgresServices "github.com/tcero76/marketplace/postgres/services"
	configRedis "github.com/tcero76/marketplace/redis/config"
	"github.com/tcero76/marketplace/redis/model"
	redisServices "github.com/tcero76/marketplace/redis/services"
)

type NullAuthCacheService struct{}

func InitDependencies() (*Dependencies, error) {

	passiveMode := os.Getenv("PASSIVE_MODE")

	log := logConfig.NewLoggerLogstash("🗄️ BFF")

	// PostgreSQL
	dbRead, err := postgresConfig.GetPostgresRead(log)
	if err != nil {
		return nil, err
	}
	dbWrite, err := postgresConfig.GetPostgresWrite(log)
	if err != nil {
		return nil, err
	}

	// Servicios de negocio
	userServices := postgresServices.NewUserService(log, dbWrite, dbRead)
	InventoryService := postgresServices.NewInventoryService(log, dbWrite, dbRead)
	postsService := postgresServices.NewPostsService(log, dbWrite, dbRead)
	jwkService := postgresServices.NewJWKService(log, dbWrite, dbRead)
	salesService := postgresServices.NewSalesService(dbRead)

	// Redis
	var authCacheService services.IAuthCacheService
	var validator oauth2.TokenValidator
	if passiveMode == "true" {
		validator = &NullTokenValidator{}
		authCacheService = &NullAuthCacheService{}
	} else {
		validator = oauth2.NewHydraTokenValidator()
		rdb := configRedis.InitRedis()
		authCacheService = redisServices.NewAuthCacheService(log, rdb)
	}

	// Hydra
	cfg := hydra.NewConfiguration()
	cfg.Servers = hydra.ServerConfigurations{
		{URL: os.Getenv("HYDRA_ADMIN_URL")},
	}
	hydraClient := hydra.NewAPIClient(cfg)
	googleClient := config.InitGoogleAuth()

	return &Dependencies{
		Log:              log,
		DBRead:           dbRead,
		DBWrite:          dbWrite,
		AuthCacheService: authCacheService,
		UserService:      userServices,
		InventoryService: InventoryService,
		PostsService:     postsService,
		JWKService:       jwkService,
		SalesService:     salesService,
		HydraClient:      hydraClient,
		Validator:        validator,
		GoogleClient:     googleClient,
	}, nil
}

type NullTokenValidator struct{}

func (v *NullTokenValidator) Validate(token string) (jwt.MapClaims, error) {
	return jwt.MapClaims{
		"sub": "123e4567-e89b-12d3-a456-426614174000",
	}, nil
}

func (s *NullAuthCacheService) SaveSession(
	key string,
	session model.SessionData,
	ctx context.Context,
) error {
	return nil
}

func (s *NullAuthCacheService) GetSession(
	key string,
	ctx context.Context,
) (*model.SessionData, error) {
	return nil, nil
}

func (s *NullAuthCacheService) StoreTokenInRedis(
	sessionID string,
	key string,
	s2 *model.FormSession,
	ctx context.Context,
) error {
	return nil
}

func (s *NullAuthCacheService) LoadTokenFromRedis(
	sessionID string,
	key string,
	ctx context.Context,
) (string, error) {
	return "", nil
}
