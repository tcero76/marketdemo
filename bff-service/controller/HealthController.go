package controller

import (
	"github.com/labstack/echo/v4"
	logConfig "github.com/tcero76/marketplace/config/log"
)

// HealthCheckHandler godoc
//
// @Summary      Health check
// @Description  Verifica el estado del servicio.
// @Tags         Health
// @Produce      json
// @Success 200 {object} map[string]string
// @Router       /health [get]
func HealthCheckHandler(log *logConfig.LoggerLogstash) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Info("HealthCheckHandler called")
		return c.JSON(200, map[string]string{
			"status": "ok",
		})
	}
}
