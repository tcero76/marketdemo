package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	logConfig "github.com/tcero76/marketplace/config/log"
)

type mockProductService struct {
	GetProductFunc  func(string) (any, error)
	GetProductsFunc func() (any, error)
}

func (m *mockProductService) GetProduct(q string) (any, error) {
	return m.GetProductFunc(q)
}

func (m *mockProductService) GetProducts() (any, error) {
	return m.GetProductsFunc()
}

func TestGetProduct_Error(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/product?product=tv", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockSvc := &mockProductService{
		GetProductFunc: func(string) (any, error) {
			return nil, errors.New("db error")
		},
	}

	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		mockSvc,
		nil,
	)

	err := ctrl.GetProduct()(c)

	if err != nil {
		t.Fatalf("handler should not return error: %v", err)
	}

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", rec.Code)
	}

	expected := "Error fetching product: db error"
	if rec.Body.String() != expected {
		t.Fatalf("unexpected body: %s", rec.Body.String())
	}
}

func TestGetProducts_OK(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockSvc := &mockProductService{
		GetProductsFunc: func() (any, error) {
			return []string{"tv", "pc"}, nil
		},
	}

	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		mockSvc,
		nil,
	)

	err := ctrl.GetProducts()(c)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestGetProducts_OK(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockSvc := &mockProductService{
		GetProductsFunc: func() (any, error) {
			return []string{"tv", "pc"}, nil
		},
	}

	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		mockSvc,
		nil,
	)

	err := ctrl.GetProducts()(c)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}
