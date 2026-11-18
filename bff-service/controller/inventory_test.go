package controller

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
	"github.com/tcero76/marketplace/bff-service/dto"
	logConfig "github.com/tcero76/marketplace/config/log"
)

type mockProductService struct {
	GetStockFunc                 func(context.Context, uuid.UUID) (dto.StockResponse, error)
	PurchaseFunc                 func(context.Context, dto.PurchaseRequest) error
	SellFunc                     func(context.Context, dto.SellRequest) error
	AdjustStockFunc              func(context.Context, dto.AdjustStockRequest) error
	DisableProductFunc           func(context.Context, uuid.UUID) error
	EnableProductFunc            func(context.Context, uuid.UUID) error
	GetProductFunc               func(string) (*dto.ProductResponse, error)
	GetProductsFunc              func() ([]dto.ProductResponse, error)
	GetSearchProductFunc         func(dto.SearchRequest) []dto.ProductResponse
	GetRecomendationsProductFunc func(context.Context, string) []dto.ProductResponse
	GetCategoriesFunc            func() ([]string, error)
}

func (m *mockProductService) GetStock(ctx context.Context, productID uuid.UUID) (dto.StockResponse, error) {
	return m.GetStockFunc(ctx, productID)
}
func (m *mockProductService) Purchase(ctx context.Context, request dto.PurchaseRequest) error {
	return m.PurchaseFunc(ctx, request)
}
func (m *mockProductService) Sell(ctx context.Context, request dto.SellRequest) error {
	return m.SellFunc(ctx, request)
}
func (m *mockProductService) AdjustStock(ctx context.Context, request dto.AdjustStockRequest) error {
	return m.AdjustStockFunc(ctx, request)
}
func (m *mockProductService) DisableProduct(ctx context.Context, productID uuid.UUID) error {
	return m.DisableProductFunc(ctx, productID)
}
func (m *mockProductService) EnableProduct(ctx context.Context, productID uuid.UUID) error {
	return m.EnableProductFunc(ctx, productID)
}

func (m *mockProductService) GetProduct(q string) (*dto.ProductResponse, error) {
	return m.GetProductFunc(q)
}

func (m *mockProductService) GetProducts() ([]dto.ProductResponse, error) {
	return m.GetProductsFunc()
}

func (m *mockProductService) GetSearchProduct(req dto.SearchRequest) []dto.ProductResponse {
	return m.GetSearchProductFunc(req)
}

func (m *mockProductService) GetRecomendationsProduct(ctx context.Context, userId string) []dto.ProductResponse {
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
func TestGetProduct(t *testing.T) {
	t.Run("Testing GetProduct Ok", func(t *testing.T) {
		c, rec := newTestContext(
			http.MethodGet,
			"/product?product=tv",
			nil,
		)
		mockSvc := &mockProductService{
			GetProductFunc: func(q string) (*dto.ProductResponse, error) {
				if q != "tv" {
					t.Fatalf("expected query %q, got %q", "tv", q)
				}
				return &dto.ProductResponse{
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
		var got dto.ProductResponse

		if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
			t.Fatalf("invalid json: %v", err)
		}
		if got.Description != "Televisor" {
			t.Fatalf("expected %q, got %q", "Televisor", got.Description)
		}
	})
	t.Run("Testing GetProduct Error", func(t *testing.T) {
		c, rec := newTestContext(
			http.MethodGet,
			"/product?product=tv",
			nil,
		)
		mockSvc := &mockProductService{
			GetProductFunc: func(string) (*dto.ProductResponse, error) {
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
		if rec.Code != http.StatusNotFound {
			t.Fatalf("expected %d, got %d", http.StatusNotFound, rec.Code)
		}
		expected := "Error fetching product: db error"
		if rec.Body.String() != expected {
			t.Fatalf("expected %q, got %q", expected, rec.Body.String())
		}
	})
	t.Run("Test GetProduct MissingQuery", func(t *testing.T) {
		c, rec := newTestContext(
			http.MethodGet,
			"/product",
			nil,
		)
		mockSvc := &mockProductService{
			GetProductFunc: func(string) (*dto.ProductResponse, error) {
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
	})
}

func TestGetProducts(t *testing.T) {
	t.Run("Test GetProducts Ok", func(t *testing.T) {
		c, rec := newTestContext(
			http.MethodGet,
			"/products",
			nil,
		)
		mockSvc := &mockProductService{
			GetProductsFunc: func() ([]dto.ProductResponse, error) {
				return []dto.ProductResponse{
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
		var products []dto.ProductResponse
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
	})
	t.Run("Test GetProducts Error", func(t *testing.T) {
		c, rec := newTestContext(
			http.MethodGet,
			"/products",
			nil,
		)
		mockSvc := &mockProductService{
			GetProductsFunc: func() ([]dto.ProductResponse, error) {
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
	})
}

func TestGetSearchProduct(t *testing.T) {
	t.Run("Test SearchProduct Ok", func(t *testing.T) {
		body := `{
		"query":"tv"
	}`
		c, rec := newTestBodyContext(
			http.MethodPost,
			"/products/search",
			strings.NewReader(body),
		)
		mockSvc := &mockProductService{
			GetSearchProductFunc: func(req dto.SearchRequest) []dto.ProductResponse {
				return []dto.ProductResponse{
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
	})
	t.Run("Test SearchProduct BadRequest", func(t *testing.T) {
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
	})
	t.Run("Test SearchProduct RequestMapped", func(t *testing.T) {
		body := `{
		"text":["televisor"]
	}`
		c, rec := newTestBodyContext(
			http.MethodPost,
			"/products/search",
			strings.NewReader(body),
		)
		var received dto.SearchRequest
		mockSvc := &mockProductService{
			GetSearchProductFunc: func(r dto.SearchRequest) []dto.ProductResponse {
				received = r
				return []dto.ProductResponse{}
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
	})
}

func TestGetRecomendationsProduct(t *testing.T) {
	t.Run("Test SearchProduct Unauthorized_NoUser", func(t *testing.T) {
		c, _ := newTestContext(
			http.MethodGet,
			"/recommendations",
			nil,
		)
		mockSvc := &mockProductService{
			GetRecomendationsProductFunc: func(ctx context.Context, userID string) []dto.ProductResponse {
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
	})
	t.Run("Test SearchProduct Unauthorized_NoSub", func(t *testing.T) {
		c, _ := newTestContext(
			http.MethodGet,
			"/recommendations",
			nil,
		)
		c.Set("user", jwt.MapClaims{
			"email": "test@test.cl",
		})
		mockSvc := &mockProductService{
			GetRecomendationsProductFunc: func(ctx context.Context, userID string) []dto.ProductResponse {
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
	})
}

func TestGetCategories(t *testing.T) {
	t.Run("Test SearchProduct Ok", func(t *testing.T) {
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
	})
	t.Run("Test SearchProduct Error", func(t *testing.T) {
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
	})
}

type mockPostsService struct {
	CreatePosteoFunc func(*dto.PostRequest, string) error
	GetPosteosFunc   func(string) ([]dto.PostResponse, error)
}

func (m *mockPostsService) GetPosts(limit int, offset int) []dto.PostResponse {
	return nil
}

func (m *mockPostsService) GetTotalPosts() int64 {
	return 0
}

func (m *mockPostsService) CreatePosteo(posteo *dto.PostRequest, userId string) error {
	return m.CreatePosteoFunc(posteo, userId)
}

func (m *mockPostsService) GetPosteos(productId string) ([]dto.PostResponse, error) {
	return m.GetPosteosFunc(productId)
}
func TestCreatePostDemo(t *testing.T) {
	t.Run("Test CreatePostDemo InvalidJSON", func(t *testing.T) {
		c, rec := newTestBodyContext(
			http.MethodPost,
			"/posteo",
			strings.NewReader("{"),
		)
		mockPosts := &mockPostsService{
			CreatePosteoFunc: func(posteo *dto.PostRequest, userId string) error {
				t.Fatal("service should not be called")
				return nil
			},
		}
		ctrl := NewPostController(
			logConfig.NewLoggerLogstash("test"),
			nil,
			mockPosts,
		)
		err := ctrl.CreatePost()(c)
		if err != nil {
			t.Fatalf("handler returned error: %v", err)
		}
		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected %d, got %d", http.StatusBadRequest, rec.Code)
		}
	})
	t.Run("Test CreatePostDemo Unauthorized_NoUser", func(t *testing.T) {
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
			CreatePosteoFunc: func(posteo *dto.PostRequest, userId string) error {
				t.Fatal("service should not be called")
				return nil
			},
		}
		ctrl := NewPostController(
			logConfig.NewLoggerLogstash("test"),
			nil,
			mockPosts,
		)
		err := ctrl.CreatePost()(c)
		httpErr, ok := err.(*echo.HTTPError)
		if !ok {
			t.Fatalf("expected HTTPError, got %T", err)
		}
		if httpErr.Code != http.StatusUnauthorized {
			t.Fatalf("expected %d, got %d", http.StatusUnauthorized, httpErr.Code)
		}
	})
	t.Run("Test CreatePostDemo Unauthorized_NoSub", func(t *testing.T) {
		c, _ := newTestBodyContext(
			http.MethodPost,
			"/posteo",
			strings.NewReader(`{"title":"demo"}`),
		)
		c.Set("user", jwt.MapClaims{
			"email": "test@test.com",
		})
		mockPosts := &mockPostsService{
			CreatePosteoFunc: func(posteo *dto.PostRequest, userId string) error {
				t.Fatal("service should not be called")
				return nil
			},
		}
		ctrl := NewPostController(
			logConfig.NewLoggerLogstash("test"),
			nil,
			mockPosts,
		)
		err := ctrl.CreatePost()(c)
		httpErr, ok := err.(*echo.HTTPError)
		if !ok {
			t.Fatalf("expected HTTPError, got %T", err)
		}
		if httpErr.Code != http.StatusUnauthorized {
			t.Fatalf("expected %d, got %d", http.StatusUnauthorized, httpErr.Code)
		}
	})
	t.Run("Test CreatePostDemo InvalidUUID", func(t *testing.T) {
		c, rec := newTestBodyContext(
			http.MethodPost,
			"/posteo",
			strings.NewReader(`{"title":"demo"}`),
		)
		c.Set("user", jwt.MapClaims{
			"sub": "no-es-un-uuid",
		})
		mockPosts := &mockPostsService{
			CreatePosteoFunc: func(posteo *dto.PostRequest, userId string) error {
				t.Fatal("service should not be called")
				return nil
			},
		}
		ctrl := NewPostController(
			logConfig.NewLoggerLogstash("test"),
			nil,
			mockPosts,
		)
		err := ctrl.CreatePost()(c)
		if err != nil {
			t.Fatalf("handler returned error: %v", err)
		}
		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected %d, got %d", http.StatusBadRequest, rec.Code)
		}
	})
	t.Run("Test CreatePostDemo CreateError", func(t *testing.T) {
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
			CreatePosteoFunc: func(posteo *dto.PostRequest, userId string) error {
				return errors.New("database error")
			},
		}
		ctrl := NewPostController(
			logConfig.NewLoggerLogstash("test"),
			nil,
			mockPosts,
		)
		err := ctrl.CreatePost()(c)
		if err != nil {
			t.Fatalf("handler returned error: %v", err)
		}
		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("expected %d, got %d", http.StatusInternalServerError, rec.Code)
		}
	})
	t.Run("Test CreatePostDemo OK", func(t *testing.T) {
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
			CreatePosteoFunc: func(posteo *dto.PostRequest, id string) error {
				called = true
				if id != userID {
					t.Fatalf("expected user %s, got %s", userID, id)
				}
				return nil
			},
		}
		ctrl := NewPostController(
			logConfig.NewLoggerLogstash("test"),
			nil,
			mockPosts,
		)
		err := ctrl.CreatePost()(c)
		if err != nil {
			t.Fatalf("handler returned error: %v", err)
		}
		if !called {
			t.Fatal("CreatePosteo was not called")
		}
		if rec.Code != http.StatusCreated {
			t.Fatalf("expected %d, got %d", http.StatusCreated, rec.Code)
		}
	})
}

func TestGetPosteosDemo(t *testing.T) {
	t.Run("Test GetPosteos OK", func(t *testing.T) {
		c, rec := newTestContext(
			http.MethodGet,
			"/posteos?productId=123",
			nil,
		)
		mockPosts := &mockPostsService{
			GetPosteosFunc: func(productId string) ([]dto.PostResponse, error) {
				if productId != "123" {
					t.Fatalf("expected productId 123, got %s", productId)
				}
				return []dto.PostResponse{
					{},
				}, nil
			},
		}
		ctrl := NewPostController(
			logConfig.NewLoggerLogstash("test"),
			nil,
			mockPosts,
		)
		err := ctrl.GetPosts()(c)
		if err != nil {
			t.Fatalf("handler returned error: %v", err)
		}
		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rec.Code)
		}
	})
	t.Run("Test GetPosteos MissingProductID", func(t *testing.T) {
		c, rec := newTestContext(
			http.MethodGet,
			"/posteos",
			nil,
		)
		mockPosts := &mockPostsService{
			GetPosteosFunc: func(productId string) ([]dto.PostResponse, error) {
				t.Fatal("service should not be called")
				return nil, nil
			},
		}
		ctrl := NewPostController(
			logConfig.NewLoggerLogstash("test"),
			nil,
			mockPosts,
		)
		err := ctrl.GetPosts()(c)
		if err != nil {
			t.Fatalf("handler returned error: %v", err)
		}

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rec.Code)
		}
	})
	t.Run("Test GetPosteos Erro", func(t *testing.T) {
		c, rec := newTestContext(
			http.MethodGet,
			"/posteos?productId=123",
			nil,
		)
		mockPosts := &mockPostsService{
			GetPosteosFunc: func(productId string) ([]dto.PostResponse, error) {
				return nil, errors.New("database error")
			},
		}
		ctrl := NewPostController(
			logConfig.NewLoggerLogstash("test"),
			nil,
			mockPosts,
		)
		err := ctrl.GetPosts()(c)
		if err != nil {
			t.Fatalf("handler returned error: %v", err)
		}
		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("expected 500, got %d", rec.Code)
		}
	})
}
