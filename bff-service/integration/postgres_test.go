package integration

import (
	"net/http/httptest"
	"strings"
	"testing"

	"encoding/json"
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/require"
	"github.com/tcero76/marketplace/bff-service/dto/demo"
	"github.com/tcero76/marketplace/bff-service/server"
	"github.com/tcero76/marketplace/postgres/model"
)

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

	var products []demo.Product
	json.Unmarshal(rec.Body.Bytes(), &products)

	require.Len(t, products, 30)
}

func TestGetProduct_Ok(t *testing.T) {
	e, err := server.StartServer()
	if err != nil {
		log.Error(err.Error())
	}
	req := httptest.NewRequest(http.MethodGet,
		"/usuario/getProduct?product=1",
		nil)
	req.Header.Set("Authorization", "Bearer cualquier-cosa")
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)
	var product demo.Product
	json.Unmarshal(rec.Body.Bytes(), &product)
	require.Equal(t, "Essence Mascara Lash Princess", product.Title)
}

func TestGetProduct_NotFound(t *testing.T) {
	e, err := server.StartServer()
	if err != nil {
		log.Error(err.Error())
	}
	req := httptest.NewRequest(http.MethodGet,
		"/usuario/getProduct?product=999",
		nil)
	req.Header.Set("Authorization", "Bearer cualquier-cosa")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetSearchProduct_Ok(t *testing.T) {
	e, err := server.StartServer()
	if err != nil {
		log.Error(err.Error())
	}
	body := `{"hashtag":"","mention":"","text":["perfect"]}`
	req := httptest.NewRequest(http.MethodPost,
		"/usuario/searchProducts",
		strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer cualquier-cosa")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)
	var products []demo.Product
	json.Unmarshal(rec.Body.Bytes(), &products)
	require.Len(t, products, 6)
}

func TestGetSearchProduct_NotFound(t *testing.T) {
	e, err := server.StartServer()
	if err != nil {
		log.Error(err.Error())
	}
	body := `{"hashtag":"","mention":"","text":["xxxxxxxx"]}`
	req := httptest.NewRequest(http.MethodPost,
		"/usuario/searchProducts",
		strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer cualquier-cosa")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)
	var products []demo.Product
	json.Unmarshal(rec.Body.Bytes(), &products)
	require.Len(t, products, 0)
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
	var categories []model.Category
	json.Unmarshal(rec.Body.Bytes(), &categories)
	require.Len(t, categories, 4)
}
