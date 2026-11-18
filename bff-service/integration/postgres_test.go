package integration

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"encoding/json"
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/require"
	"github.com/tcero76/marketplace/bff-service/dto"
	"github.com/tcero76/marketplace/bff-service/server"
	inventoryModel "github.com/tcero76/marketplace/postgres/model/inventory"
)

func TestGetProduct(t *testing.T) {
	e, err := server.StartServer()
	require.NoError(t, err)
	t.Run("Test GetProduct Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet,
			"/usuario/getProduct?product=831ad86a-6956-475f-8a61-5a89d72b130d",
			nil)
		req.Header.Set("Authorization", "Bearer cualquier-cosa")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
		var product dto.ProductResponse
		json.Unmarshal(rec.Body.Bytes(), &product)
		require.Equal(t, "Red Lipstick", product.Title)
	})
	t.Run("Test GetProduct NotFound", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet,
			"/usuario/getProduct?product=45a78d3b-55f2-4d28-ba6a-d25e75749055",
			nil)
		req.Header.Set("Authorization", "Bearer cualquier-cosa")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		require.Equal(t, http.StatusNotFound, rec.Code)

	})
}

func TestGetProducts(t *testing.T) {
	e, err := server.StartServer()
	if err != nil {
		log.Error(err.Error())
	}
	req := httptest.NewRequest(http.MethodGet,
		"/usuario/getProducts",
		nil)
	req.Header.Set("Authorization", "Bearer cualquier-cosa")
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)

	var products []dto.ProductResponse
	json.Unmarshal(rec.Body.Bytes(), &products)
	fmt.Print("productos: ", products)
	require.Len(t, products, 30)
}

func TestGetSearchProduct(t *testing.T) {
	e, err := server.StartServer()
	if err != nil {
		log.Error(err.Error())
	}
	t.Run("Test GetSearchProduct Ok", func(t *testing.T) {
		body := `{"hashtag":"","mention":"","text":["perfect"]}`
		req := httptest.NewRequest(http.MethodPost,
			"/usuario/searchProducts",
			strings.NewReader(body))
		req.Header.Set("Authorization", "Bearer cualquier-cosa")
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
		var products []dto.ProductResponse
		json.Unmarshal(rec.Body.Bytes(), &products)
		require.Len(t, products, 6)
	})
	t.Run("Test GetSearchProduct NotFound", func(t *testing.T) {
		body := `{"hashtag":"","mention":"","text":["xxxxxxxx"]}`
		req := httptest.NewRequest(http.MethodPost,
			"/usuario/searchProducts",
			strings.NewReader(body))
		req.Header.Set("Authorization", "Bearer cualquier-cosa")
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
		var products []dto.ProductResponse
		json.Unmarshal(rec.Body.Bytes(), &products)
		require.Len(t, products, 0)
	})
}

func TestGetCategories_Ok(t *testing.T) {
	e, err := server.StartServer()
	if err != nil {
		log.Error(err.Error())
	}
	req := httptest.NewRequest(
		http.MethodGet,
		"/usuario/getCategories",
		nil)
	req.Header.Set("Authorization", "Bearer cualquier-cosa")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)
	var categories []inventoryModel.Category
	json.Unmarshal(rec.Body.Bytes(), &categories)
	require.Len(t, categories, 4)
}
