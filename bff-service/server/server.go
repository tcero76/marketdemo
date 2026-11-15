package server

import (
	"context"
	"os"

	echoSwagger "github.com/swaggo/echo-swagger"
	formLLM "github.com/tcero76/marketplace/bff-service/bot/llm"
	"github.com/tcero76/marketplace/bff-service/config"
	"github.com/tcero76/marketplace/bff-service/controller"
	_ "github.com/tcero76/marketplace/bff-service/docs"
	"github.com/tcero76/marketplace/bff-service/oauth2/cor"
	logConfig "github.com/tcero76/marketplace/config/log"

	"github.com/tcero76/marketplace/config/metrics"
	modelServices "github.com/tcero76/marketplace/postgres/services"

	"github.com/tcero76/marketplace/bff-service/oauth2"

	"github.com/labstack/echo/v4"
	hydra "github.com/ory/hydra-client-go/v2"
)

var hydraAdminClient *hydra.APIClient

var ctx = context.Background()

func StartServer() (*echo.Echo, error) {

	e := echo.New()

	metrics.Register()
	e.Use(metrics.Middleware())
	metrics.RegisterEndpoint(e)

	log := logConfig.NewLoggerLogstash("🗄️  BFF")
	log.Info("Iniciando servidor...")

	d, err := InitDependencies()
	if err != nil {
		log.Error("Error de inicialización")
		return nil, err
	}

	internalAuth := cor.NewInternalAuth(d.HydraClient, d.UserService)
	googleAuth := cor.NewGoogleAuth(d.GoogleClient, d.UserService)
	loginHandler, consentHandler, callbackHandler :=
		oauth2.InitOauth2Handlers(d.AuthCacheService, d.UserService, internalAuth, googleAuth)

	e.GET("/health", controller.HealthCheckHandler(log))
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	session := e.Group("")
	session.Use(config.RedisSessionMiddleware(d.AuthCacheService, log))
	session.Use(config.LoggerMiddleware(log))

	authController := controller.NewAuthController(log)

	session.POST("/login", authController.HandleLogin(loginHandler))
	session.GET("/consent", authController.HandleConsent(consentHandler))
	session.GET("/callback", authController.HandleCallback(callbackHandler))
	session.GET("/getAuthentication", authController.AuthHandler())
	session.POST("/refresh", authController.RefreshTokenHandler(d.AuthCacheService))
	session.GET("/logout", authController.LogoutHandler(d.AuthCacheService))
	session.POST("/signup", authController.SignUpHandler(d.UserService))

	protegido := session.Group("/usuario")
	protegido.Use(oauth2.JWTMiddleware(log, d.Validator))

	productController := controller.NewProductController(log, d.ProductService, d.PostsService)
	fileController := controller.NewFileController(log)
	jwkController := controller.NewJwksHandler(log, d.JWKService)

	protegido.GET("/getProduct", productController.GetProduct())
	protegido.GET("/getProducts", productController.GetProducts())
	protegido.POST("/searchProducts", productController.GetSearchProduct())
	protegido.GET("/getCategories", productController.GetCategories())
	protegido.GET("/product/getRecommendations", productController.GetRecomendationsProduct())

	protegido.POST("/createPostDemo", productController.CreatePosteoDemo())
	protegido.GET("/getPosteosDemo", productController.GetPosteosDemo())

	chatboxServices := modelServices.NewService(d.DBWrite)
	processorLLM := formLLM.NewBotLLM(d.AuthCacheService, chatboxServices)
	protegido.POST("/chat", controller.ChatLLMHandler(processorLLM))

	salesController := controller.NewSalesController(d.SalesService, log)
	protegido.GET("/getOrder", salesController.GetOrder())

	e.GET("/webhook", controller.WebhookHandler())
	e.POST("/webhook", controller.WebhookHandler())

	e.POST("/uploadImage", fileController.UploadImage())
	e.GET("/getImage/:name", fileController.GetImage())

	e.GET("/.well-known/jwks.json", jwkController.JwksHandler())
	e.POST("/token", jwkController.TokenHandler())

	e.GET("/embeded", controller.GetEmbeded(log))
	log.Info("Servidor iniciado en el puerto: ", os.Getenv("PORT"))
	return e, nil
}
