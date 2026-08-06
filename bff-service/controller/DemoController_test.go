package controller

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"strings"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"io"

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

func newTestContext(
	method string,
	url string,
	body io.Reader,
) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(
		method,
		url,
		body,
	)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func newTestBodyContext(
	method string,
	url string,
	body io.Reader,
) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(
		method,
		url,
		body,
	)
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return e.NewContext(req, rec), rec
}

func TestGetProduct_Error(t *testing.T) {
	c, rec := newTestContext(
		http.MethodGet,
		"/product?product=tv",
		nil,
	)
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
	c, rec := newTestContext(
		http.MethodGet,
		"/product?product=tv",
		nil,
	)
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
	c, rec := newTestContext(
		http.MethodGet,
		"/product",
		nil,
	)
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
	c, rec := newTestContext(
		http.MethodGet,
		"/products",
		nil,
	)
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
	c, rec := newTestContext(
		http.MethodGet,
		"/products",
		nil,
	)
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

func TestGetSearchProduct_BadRequest(t *testing.T) {
	c, rec := newTestContext(
		http.MethodGet,
		"/products/search",
		strings.NewReader("{"),
	)
	mockSvc := &mockProductService{}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		mockSvc,
		nil,
	)
	err := ctrl.GetSearchProduct()(c)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestGetSearchProduct_OK(t *testing.T) {
	body := `{
		"query":"tv"
	}`
	c, rec := newTestBodyContext(
		http.MethodPost,
		"/products/search",
		strings.NewReader(body),
	)
	mockSvc := &mockProductService{
		GetSearchProductFunc: func(req payload.SearchRequest) []demo.Product {
			return []demo.Product{
				{
					Description: "Televisor",
				},
			}
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		mockSvc,
		nil,
	)
	err := ctrl.GetSearchProduct()(c)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestGetSearchProduct_RequestMapped(t *testing.T) {
	body := `{
		"text":["televisor"]
	}`
	c, rec := newTestBodyContext(
		http.MethodPost,
		"/products/search",
		strings.NewReader(body),
	)
	var received payload.SearchRequest
	mockSvc := &mockProductService{
		GetSearchProductFunc: func(r payload.SearchRequest) []demo.Product {
			received = r
			return []demo.Product{}
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		mockSvc,
		nil,
	)
	err := ctrl.GetSearchProduct()(c)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rec.Code)
	}
	if received.Text[0] != "televisor" {
		t.Fatalf("expected text 'televisor', got '%s'", received.Text)
	}
}

func TestGetRecomendationsProduct_Unauthorized_NoUser(t *testing.T) {
	c, _ := newTestContext(
		http.MethodGet,
		"/recommendations",
		nil,
	)
	mockSvc := &mockProductService{
		GetRecomendationsProductFunc: func(ctx context.Context, userID string) []int {
			t.Fatal("service should not be called")
			return nil
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		mockSvc,
		nil,
	)
	err := ctrl.GetRecomendationsProduct()(c)
	if err == nil {
		t.Fatal("expected an HTTP error")
	}
	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("expected *echo.HTTPError, got %T", err)
	}
	if httpErr.Code != http.StatusUnauthorized {
		t.Fatalf("expected %d, got %d", http.StatusUnauthorized, httpErr.Code)
	}
}

func TestGetRecomendationsProduct_Unauthorized_NoSub(t *testing.T) {
	c, _ := newTestContext(
		http.MethodGet,
		"/recommendations",
		nil,
	)
	c.Set("user", jwt.MapClaims{
		"email": "test@test.cl",
	})
	mockSvc := &mockProductService{
		GetRecomendationsProductFunc: func(ctx context.Context, userID string) []int {
			t.Fatal("service should not be called")
			return nil
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		mockSvc,
		nil,
	)
	err := ctrl.GetRecomendationsProduct()(c)
	if err == nil {
		t.Fatal("expected an HTTP error")
	}
	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("expected *echo.HTTPError, got %T", err)
	}
	if httpErr.Code != http.StatusUnauthorized {
		t.Fatalf("expected %d, got %d", http.StatusUnauthorized, httpErr.Code)
	}
}

func TestGetCategories_OK(t *testing.T) {
	c, rec := newTestContext(
		http.MethodGet,
		"/categories",
		nil,
	)
	mockSvc := &mockProductService{
		GetCategoriesFunc: func() ([]string, error) {
			return []string{"TV", "Computación", "Celulares"}, nil
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		mockSvc,
		nil,
	)
	err := ctrl.GetCategories()(c)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rec.Code)
	}
	expected := `["TV","Computación","Celulares"]`
	if strings.TrimSpace(rec.Body.String()) != expected {
		t.Fatalf("expected %s, got %s", expected, rec.Body.String())
	}
}

func TestGetCategories_Error(t *testing.T) {
	c, rec := newTestContext(
		http.MethodPost,
		"/categories",
		nil,
	)
	mockSvc := &mockProductService{
		GetCategoriesFunc: func() ([]string, error) {
			return nil, errors.New("database error")
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		mockSvc,
		nil,
	)
	err := ctrl.GetCategories()(c)

	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d, got %d", http.StatusInternalServerError, rec.Code)
	}

	expected := `{"error":"Error fetching categories"}`

	if strings.TrimSpace(rec.Body.String()) != expected {
		t.Fatalf("expected %s, got %s", expected, rec.Body.String())
	}
}

type mockPostsService struct {
	CreatePosteoFunc func(*demo.PostDTO, string) error
	GetPosteosFunc   func(string) ([]demo.PostDTO, error)
}

func (m *mockPostsService) GetPosts(limit int, offset int) []demo.PostDTO {
	return nil
}

func (m *mockPostsService) GetTotalPosts() int64 {
	return 0
}

func (m *mockPostsService) CreatePosteo(posteo *demo.PostDTO, userId string) error {
	return m.CreatePosteoFunc(posteo, userId)
}

func (m *mockPostsService) GetPosteos(productId string) ([]demo.PostDTO, error){
	return m.GetPosteosFunc(productId)
}

func TestCreatePosteoDemo_BadRequest_InvalidJSON(t *testing.T) {
	c, rec := newTestBodyContext(
		http.MethodPost,
		"/posteo",
		strings.NewReader("{"),
	)
	mockPosts := &mockPostsService{
		CreatePosteoFunc: func(posteo *demo.PostDTO, userId string) error {
			t.Fatal("service should not be called")
			return nil
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		nil,
		mockPosts,
	)
	err := ctrl.CreatePosteoDemo()(c)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestCreatePosteoDemo_Unauthorized_NoUser(t *testing.T) {
	body := `{
		"title":"Mi post",
		"description":"demo"
	}`
	c, _ := newTestBodyContext(
		http.MethodPost,
		"/posteo",
		strings.NewReader(body),
	)
	mockPosts := &mockPostsService{
		CreatePosteoFunc: func(posteo *demo.PostDTO, userId string) error {
			t.Fatal("service should not be called")
			return nil
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		nil,
		mockPosts,
	)
	err := ctrl.CreatePosteoDemo()(c)
	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("expected HTTPError, got %T", err)
	}
	if httpErr.Code != http.StatusUnauthorized {
		t.Fatalf("expected %d, got %d", http.StatusUnauthorized, httpErr.Code)
	}
}

func TestCreatePosteoDemo_Unauthorized_NoSub(t *testing.T) {
	c, _ := newTestBodyContext(
		http.MethodPost,
		"/posteo",
		strings.NewReader(`{"title":"demo"}`),
	)
	c.Set("user", jwt.MapClaims{
		"email": "test@test.com",
	})
	mockPosts := &mockPostsService{
		CreatePosteoFunc: func(posteo *demo.PostDTO, userId string) error {
			t.Fatal("service should not be called")
			return nil
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		nil,
		mockPosts,
	)
	err := ctrl.CreatePosteoDemo()(c)
	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("expected HTTPError, got %T", err)
	}
	if httpErr.Code != http.StatusUnauthorized {
		t.Fatalf("expected %d, got %d", http.StatusUnauthorized, httpErr.Code)
	}
}

func TestCreatePosteoDemo_InvalidUUID(t *testing.T) {
	c, rec := newTestBodyContext(
		http.MethodPost,
		"/posteo",
		strings.NewReader(`{"title":"demo"}`),
	)
	c.Set("user", jwt.MapClaims{
		"sub": "no-es-un-uuid",
	})
	mockPosts := &mockPostsService{
		CreatePosteoFunc: func(posteo *demo.PostDTO, userId string) error {
			t.Fatal("service should not be called")
			return nil
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		nil,
		mockPosts,
	)
	err := ctrl.CreatePosteoDemo()(c)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestCreatePosteoDemo_CreateError(t *testing.T) {
	c, rec := newTestBodyContext(
		http.MethodPost,
		"/posteo",
		strings.NewReader(`{"title":"demo"}`),
	)
	userID := uuid.New().String()
	c.Set("user", jwt.MapClaims{
		"sub": userID,
	})
	mockPosts := &mockPostsService{
		CreatePosteoFunc: func(posteo *demo.PostDTO, userId string) error {
			return errors.New("database error")
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		nil,
		mockPosts,
	)
	err := ctrl.CreatePosteoDemo()(c)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d, got %d", http.StatusInternalServerError, rec.Code)
	}
}

func TestCreatePosteoDemo_OK(t *testing.T) {
	c, rec := newTestBodyContext(
		http.MethodPost,
		"/posteo",
		strings.NewReader(`{"title":"demo"}`),
	)
	userID := uuid.New().String()
	c.Set("user", jwt.MapClaims{
		"sub": userID,
	})
	called := false
	mockPosts := &mockPostsService{
		CreatePosteoFunc: func(posteo *demo.PostDTO, id string) error {
			called = true
			if id != userID {
				t.Fatalf("expected user %s, got %s", userID, id)
			}
			return nil
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		nil,
		mockPosts,
	)
	err := ctrl.CreatePosteoDemo()(c)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if !called {
		t.Fatal("CreatePosteo was not called")
	}
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected %d, got %d", http.StatusCreated, rec.Code)
	}
}

func TestGetPosteosDemo_OK(t *testing.T) {
	c, rec := newTestContext(
		http.MethodGet,
		"/posteos?productId=123",
		nil,
	)
	mockPosts := &mockPostsService{
		GetPosteosFunc: func(productId string) ([]demo.PostDTO, error) {
			if productId != "123" {
				t.Fatalf("expected productId 123, got %s", productId)
			}
			return []demo.PostDTO{
				{},
			}, nil
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		nil,
		mockPosts,
	)
	err := ctrl.GetPosteosDemo()(c)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestGetPosteosDemo_MissingProductID(t *testing.T) {
	c, rec := newTestContext(
		http.MethodGet,
		"/posteos",
		nil,
	)
	mockPosts := &mockPostsService{
		GetPosteosFunc: func(productId string) ([]demo.PostDTO, error) {
			t.Fatal("service should not be called")
			return nil, nil
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		nil,
		mockPosts,
	)
	err := ctrl.GetPosteosDemo()(c)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
}

func TestGetPosteosDemo_Error(t *testing.T) {
	c, rec := newTestContext(
		http.MethodGet,
		"/posteos?productId=123",
		nil,
	)
	mockPosts := &mockPostsService{
		GetPosteosFunc: func(productId string) ([]demo.PostDTO, error) {
			return nil, errors.New("database error")
		},
	}
	ctrl := NewProductController(
		logConfig.NewLoggerLogstash("test"),
		nil,
		mockPosts,
	)
	err := ctrl.GetPosteosDemo()(c)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", rec.Code)
	}
}