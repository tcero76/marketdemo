package controller

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"

	"github.com/labstack/echo/v4"
	logConfig "github.com/tcero76/marketplace/config/log"
	"github.com/tcero76/marketplace/bff-service/dto/demo"
	"github.com/tcero76/marketplace/bff-service/payload"
)

type mockProductService struct {
	GetProductFunc                func(string) (*demo.Product, error)
	GetProductsFunc               func() ([]demo.Product, error)
	GetSearchProductFunc          func(payload.SearchRequest) []demo.Product
	GetRecomendationsProductFunc  func(context.Context, string) []int
	GetCategoriesFunc             func() ([]string, error)
}

func (m *mockProductService) GetProduct(q string) (*demo.Product, error) {
	return m.GetProductFunc(q)
}

func (m *mockProductService) GetProducts() ([]demo.Product, error) {
	return m.GetProductsFunc()
}

func (m *mockProductService) GetSearchProduct(req payload.SearchRequest) []demo.Product {
	return m.GetSearchProductFunc(req)
}

func (m *mockProductService) GetRecomendationsProduct(ctx context.Context, userId string) []int {
	return m.GetRecomendationsProductFunc(ctx, userId)
}

func (m *mockProductService) GetCategories() ([]string, error) {
	return m.GetCategoriesFunc()
}

func TestGetProduct_Error(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/product?product=tv", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockSvc := &mockProductService{
		GetProductFunc: func(string) (*demo.Product, error) {
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
		t.Fatalf("handler returned error: %v", err)
	}
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d, got %d", http.StatusInternalServerError, rec.Code)
	}
	expected := "Error fetching product: db error"
	if rec.Body.String() != expected {
		t.Fatalf("expected %q, got %q", expected, rec.Body.String())
	}
}

func TestGetProduct_OK(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/product?product=tv", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockSvc := &mockProductService{
		GetProductFunc: func(q string) (*demo.Product, error) {
			if q != "tv" {
				t.Fatalf("expected query %q, got %q", "tv", q)
			}
			return &demo.Product{
				Description: "Televisor",
			}, nil
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		mockSvc,
		nil,
	)
	err := ctrl.GetProduct()(c)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rec.Code)
	}
	var got demo.Product

	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("invalid json: %v", err)
	}

	if got.Description != "Televisor" {
		t.Fatalf("expected %q, got %q", "Televisor", got.Description)
	}
}

func TestGetProduct_MissingQuery(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/product", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockSvc := &mockProductService{
		GetProductFunc: func(string) (*demo.Product, error) {
			t.Fatal("GetProduct should not have been called")
			return nil, nil
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		mockSvc,
		nil,
	)
	err := ctrl.GetProduct()(c)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, rec.Code)
	}
	expected := "missing product parameter"
	if rec.Body.String() != expected {
		t.Fatalf("expected %q, got %q", expected, rec.Body.String())
	}
}

func TestGetProducts_OK(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockSvc := &mockProductService{
		GetProductsFunc: func() ([]demo.Product, error) {
			return []demo.Product{
				{Description: "Televisor"},
				{Description: "Notebook"},
			}, nil
		},
	}

	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		mockSvc,
		nil,
	)

	err := ctrl.GetProducts()(c)

	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rec.Code)
	}

	var products []demo.Product

	err = json.Unmarshal(rec.Body.Bytes(), &products)
	if err != nil {
		t.Fatalf("invalid json: %v", err)
	}

	if len(products) != 2 {
		t.Fatalf("expected 2 products, got %d", len(products))
	}

	if products[0].Description != "Televisor" {
		t.Fatalf("expected first product %q, got %q", "Televisor", products[0].Description)
	}

	if products[1].Description != "Notebook" {
		t.Fatalf("expected second product %q, got %q", "Notebook", products[1].Description)
	}
}

func TestGetProducts_Error(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockSvc := &mockProductService{
		GetProductsFunc: func() ([]demo.Product, error) {
			return nil, errors.New("db error")
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		mockSvc,
		nil,
	)
	err := ctrl.GetProducts()(c)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d, got %d", http.StatusInternalServerError, rec.Code)
	}
	var message string
	err = json.Unmarshal(rec.Body.Bytes(), &message)
	if err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	expected := "Error fetching GetProducts"
	if message != expected {
		t.Fatalf("expected %q, got %q", expected, message)
	}
}