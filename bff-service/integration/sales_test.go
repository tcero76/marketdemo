package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/require"
	"github.com/tcero76/marketplace/bff-service/dto/sales"
	"github.com/tcero76/marketplace/bff-service/server"
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
		var order sales.GetOrderResponse
		json.Unmarshal(rec.Body.Bytes(), &order)
		require.Equal(t, nil, order)
	})
}
