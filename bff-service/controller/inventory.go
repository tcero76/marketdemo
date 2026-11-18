package controller

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/tcero76/marketplace/bff-service/dto"
	"github.com/tcero76/marketplace/bff-service/services"
	logger "github.com/tcero76/marketplace/config/log"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	log              *logger.LoggerLogstash
	inventoryService services.IInventory
	postsService     services.IDemoPosts
}

func NewProductController(log *logger.LoggerLogstash,
	inventoryService services.IInventory,
	postsService services.IDemoPosts) *ProductController {
	return &ProductController{log, inventoryService, postsService}
}

// GetProduct godoc
//
// @Summary      Obtener un producto
// @Description  Busca un producto por nombre.
// @Tags         Products
// @Produce      json
// @Param        product query string true "Nombre del producto"
// @Success      200 {object} demo.Product
// @Failure      400 {string} string "Missing product parameter"
// @Failure      404 {string} string "Error fetching product"
// @Security	  BearerAuth
// @Router       /usuario/getProduct [get]
func (h *ProductController) GetProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info("Entrando a GetProduct")
		query := c.QueryParam("product")
		if query == "" {
			return c.String(http.StatusBadRequest, "missing product parameter")
		}
		h.log.Debug("Query parametro product: ", query)
		product, err := h.inventoryService.GetProduct(query)
		if err != nil {
			h.log.Error("Error in GetProduct: ", err)
			return c.String(http.StatusNotFound, "Error fetching product: "+err.Error())
		}
		h.log.Debug("Product found: ", product)
		return c.JSON(http.StatusOK, product)
	}
}

// GetProducts godoc
//
// @Summary      Obtener productos
// @Description  Obtiene el catálogo de productos.
// @Tags         Products
// @Produce      json
// @Success      200 {array} demo.Product
// @Failure      500 {string} string "Internal server error"
// @Security     BearerAuth
// @Router       /usuario/getProducts [get]
func (h *ProductController) GetProducts() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info("Entrando a GetProducts")
		products, err := h.inventoryService.GetProducts()
		if err != nil {
			h.log.Error("Error in GetProducts: ", err)
			return c.JSON(http.StatusInternalServerError, "Error fetching GetProducts")
		}
		h.log.Debug("Products found: ", products)
		return c.JSON(http.StatusOK, products)
	}
}

// GetSearchProduct godoc
//
// @Summary      Buscar productos
// @Description  Busca productos según los criterios proporcionados.
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        request body payload.SearchRequest true "Criterios de búsqueda"
// @Success      200 {array} demo.Product
// @Failure      400 {object} map[string]string
// @Router       /products/search [post]
func (h *ProductController) GetSearchProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info("Entrando a search Products")
		var req dto.SearchRequest
		if err := c.Bind(&req); err != nil {
			h.log.Error("Error al parsear JSON: ", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Error al parsear JSON",
			})
		}
		h.log.Debug("Request de search: ", req)
		modeloSearchs := h.inventoryService.GetSearchProduct(req)
		h.log.Debug("Searchs encontrados: ", modeloSearchs)
		return c.JSON(http.StatusOK, modeloSearchs)
	}
}

// GetRecomendationsProduct godoc
//
// @Summary      Obtener recomendaciones de productos
// @Description  Obtiene recomendaciones personalizadas para el usuario autenticado.
// @Tags         Products
// @Produce      json
// @Success      200 {array} demo.Product
// @Failure      401 {object} map[string]string "Unauthorized"
// @Security     BearerAuth
// @Router       /usuario/product/getRecommendations [get]
func (h *ProductController) GetRecomendationsProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info("Fetching recommendations for user")
		user, ok := c.Get("user").(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		userId, ok := user["sub"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		h.log.Debug("User ID: ", userId)
		items := h.inventoryService.GetRecomendationsProduct(c.Request().Context(), userId)
		return c.JSON(http.StatusOK, items)
	}
}

// GetCategories godoc
//
// @Summary      Obtener categorías
// @Description  Obtiene las categorías disponibles.
// @Tags         Categories
// @Produce      json
// @Success      200 {array} string
// @Failure      500 {object} map[string]string
// @Router       /categories [get]
func (h *ProductController) GetCategories() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info("Fetching categories")
		categories, err := h.inventoryService.GetCategories()
		if err != nil {
			h.log.Error("Error fetching categories: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Error fetching categories",
			})
		}
		h.log.Debug("Categories found: ", categories)
		return c.JSON(http.StatusOK, categories)
	}
}
