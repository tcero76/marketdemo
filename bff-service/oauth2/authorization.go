package oauth2

import (
	"net/http"
	"strings"

	logger "github.com/tcero76/marketplace/config/log"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(log *logger.LoggerLogstash, validator TokenValidator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")
			if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
				log.Warn("Missing or invalid Authorization header")
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid Authorization header")
			}
			token := strings.TrimPrefix(auth, "Bearer ")
            claims, err := validator.Validate(token)
			if err != nil {
				log.Warn("Invalid token")
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			c.Set("user", claims)
			return next(c)
		}
	}
}
