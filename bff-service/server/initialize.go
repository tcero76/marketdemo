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
	modelServices "github.com/tcero76/marketplace/postgres/services"
	demoServices "github.com/tcero76/marketplace/postgres/services/demo"
	salesServices "github.com/tcero76/marketplace/postgres/services/sales"
	configRedis "github.com/tcero76/marketplace/redis/config"
	"github.com/tcero76/marketplace/redis/model"
	redisServices "github.com/tcero76/marketplace/redis/services"
)

type NullAuthCacheService struct{}

func InitDependencies() (*Dependencies, error) {

	passiveMode := os.Getenv("PASSIVE_MODE")

	log := logConfig.NewLoggerLogstash("🗄️ BFF")

	// PostgreSQL
	dbRead := postgresConfig.GetPostgresRead(log)
	dbWrite := postgresConfig.GetPostgresWrite(log)

	// Servicios de negocio
	userServices := modelServices.NewUserService(log, dbWrite, dbRead)
	productService := demoServices.NewProductService(log, dbWrite, dbRead)
	postsService := demoServices.NewPostsService(log, dbWrite, dbRead)
	jwkService := modelServices.NewJWKService(log, dbWrite, dbRead)
	salesService := salesServices.NewSalesService(dbRead)

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
		ProductService:   productService,
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
		"sub": "integration-user",
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
