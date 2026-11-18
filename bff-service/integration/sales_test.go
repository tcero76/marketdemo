package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tcero76/marketplace/bff-service/dto"
	"github.com/tcero76/marketplace/bff-service/server"
	logConfig "github.com/tcero76/marketplace/config/log"
	"github.com/tcero76/marketplace/postgres/config"
)

func TestGetOrder(t *testing.T) {
	e, err := server.StartServer()
	require.NoError(t, err)
	t.Run("Test GetOrder Success", func(t *testing.T) {
		log.Print("Entra a Test GetOrder Success")
		req := httptest.NewRequest(http.MethodGet,
			"/usuario/getOrder?order_id=1",
			nil)
		req.Header.Set("Authorization", "Bearer cualquier-cosa")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
		var order dto.GetOrderResponse
		json.Unmarshal(rec.Body.Bytes(), &order)
		require.Equal(t, nil, order)
	})
}

func TestCreateOrder(t *testing.T) {
	e, err := server.StartServer()
	assert.NoError(t, err)
	log := logConfig.NewLoggerLogstash("🗄️ BFFTesting")
	db, err := config.GetPostgresWrite(log)
	assert.NoError(t, err)
	t.Run("Test CreateOrder Success", func(t *testing.T) {
		log.Info("Entra a Test CreateOrder Success")
		body := `{
		"user_id": "123e4567-e89b-12d3-a456-426614174000",
		"currency": "CLP",
		"items": [
			{
			"product_id": "2657a1e2-38c5-4d71-9d7b-94664f66189b",
			"quantity": 2
			},
			{
			"product_id": "831ad86a-6956-475f-8a61-5a89d72b130d",
			"quantity": 1
			}
		]
		}`
		req := httptest.NewRequest(
			http.MethodPost,
			"/usuario/createOrder",
			strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer cualquier-cosa")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusCreated, rec.Code)
		var res dto.CreateOrderResponse
		json.Unmarshal(rec.Body.Bytes(), &res)
		assert.Equal(t, "PENDING", res.Status)
		deleteOrder(t, db, res.OrderID)
	})
	t.Run("Test CreateOrder UserNotFound", func(t *testing.T) {
		log.Info("Entra a Test CreateOrder UserNotFound")
		userID := "00000000-0000-0000-0000-000000004000"
		body := `{
			"user_id": "` + userID + `",
			"currency": "ECL"
			}`
		req := httptest.NewRequest(
			http.MethodPost,
			"/usuario/createOrder",
			strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer cualquier-cosa")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("Test CreateOrder InvalidUser", func(t *testing.T) {
		log.Info("Entra a Test CreateOrder InvalidUser")
		body := `{
			"user_id": "",
			"currency": "ECL"
			}`
		req := httptest.NewRequest(
			http.MethodPost,
			"/usuario/createOrder",
			strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer cualquier-cosa")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
