package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	dto "github.com/tcero76/marketplace/bff-service/dto/sales"
	"github.com/tcero76/marketplace/bff-service/mock"
	logConfig "github.com/tcero76/marketplace/config/log"
	"go.uber.org/mock/gomock"
)

func TestGetOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSales := mock.NewMockISales(ctrl)
	e := echo.New()
	log := logConfig.NewLoggerLogstash("🗄️ BFF")
	controller := &SalesController{
		log:          log,
		SalesService: mockSales,
	}
	ID := uuid.New()
	reqDTO := dto.GetOrderRequest{
		OrderID: ID,
	}
	expected := dto.GetOrderResponse{
		ID:          ID,
		UserID:      uuid.New(),
		Status:      "Ok",
		TotalAmount: decimal.NewFromInt(1),
		Currency:    "CL",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		// Items []OrderItemResponse `json:"items"`
	}
	t.Run("Test GetOrder Success", func(t *testing.T) {
		mockSales.EXPECT().
			GetOrder(gomock.Any(), gomock.Eq(reqDTO)).
			Return(expected, nil).
			Times(1)
		req := httptest.NewRequest(
			http.MethodGet,
			"/order?order_id="+ID.String(),
			nil,
		)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		controller.GetOrder()(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		var response dto.GetOrderResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.True(t, expected.CreatedAt.Equal(response.CreatedAt))
		assert.True(t, expected.UpdatedAt.Equal(response.UpdatedAt))
		assert.Equal(t, expected.ID, response.ID)
		assert.Equal(t, expected.UserID, response.UserID)
		assert.Equal(t, expected.Status, response.Status)
		assert.Equal(t, expected.TotalAmount, response.TotalAmount)
	})
	t.Run("Test GetOrder Error parser JSON", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/order?idorder="+ID.String(),
			nil,
		)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		controller.GetOrder()(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
