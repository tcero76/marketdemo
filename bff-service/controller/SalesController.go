package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tcero76/marketplace/bff-service/dto/sales"
	"github.com/tcero76/marketplace/bff-service/services"
	logger "github.com/tcero76/marketplace/config/log"
)

type SalesController struct {
	SalesService services.ISales
	log          *logger.LoggerLogstash
}

func NewSalesController(salesService services.ISales, log *logger.LoggerLogstash) *SalesController {
	return &SalesController{SalesService: salesService, log: log}
}

func (h *SalesController) GetOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info("Entrando a GetOrder")
		var req sales.GetOrderRequest
		if err := c.Bind(&req); err != nil {
			h.log.Error("Error al parsear JSON: ", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Error al parsear JSON",
			})
		}
		if req.OrderID == uuid.Nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "order_id es obligatorio",
			})
		}
		res, err := h.SalesService.GetOrder(c.Request().Context(), req)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return err
		}
		c.JSON(http.StatusOK, res)
		return nil
	}
}
