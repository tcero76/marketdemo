package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"

	"github.com/labstack/echo/v4"
	logConfig "github.com/tcero76/marketplace/config/log"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func TestHealthCheckHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	logger := logConfig.NewLoggerLogstash("🗄️ BFF")
	handler := HealthCheckHandler(logger)
	err := handler(ctx)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rec.Code)
	}
	var resp HealthResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if resp.Status != "ok" {
		t.Fatalf("expected status %q, got %q", "ok", resp.Status)
	}
}