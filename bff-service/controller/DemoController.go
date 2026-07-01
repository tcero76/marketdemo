package controller

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/tcero76/marketplace/bff-service/dto/demo"
	"github.com/tcero76/marketplace/bff-service/payload"
	"github.com/tcero76/marketplace/bff-service/services"
	logger "github.com/tcero76/marketplace/config/log"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	log            *logger.LoggerLogstash
	productService services.IDemoService
	postsService   services.IDemoPosts
}

func NewProductController(log *logger.LoggerLogstash,
	productService services.IDemoService,
	postsService services.IDemoPosts) *ProductController {
	return &ProductController{log, productService, postsService}
}

func (h *ProductController) GetProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info("Entrando a GetProduct")
		query := c.QueryParam("product")
		h.log.Debug("Query parametro product: ", query)
		product, err := h.productService.GetProduct(query)
		if err != nil {
			h.log.Error("Error in GetProduct: ", err)
			return c.String(http.StatusInternalServerError, "Error fetching product: "+err.Error())
		}

		h.log.Debug("Product found: ", product)
		return c.JSON(http.StatusOK, product)
	}
}

func (h *ProductController) GetProducts() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info("Entrando a GetProducts")
		products, err := h.productService.GetProducts()
		if err != nil {
			h.log.Error("Error in GetProducts")
			c.JSON(http.StatusInternalServerError, "Error fetching GetProducts")
		}
		h.log.Debug("Products found: ", products)
		return c.JSON(http.StatusOK, products)
	}
}

func (h *ProductController) GetSearchProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info("Entrando a search Products")
		var req payload.SearchRequest
		if err := c.Bind(&req); err != nil {
			h.log.Error("Error al parsear JSON: ", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Error al parsear JSON",
			})
		}
		h.log.Debug("Request de search: ", req)
		modeloSearchs := h.productService.GetSearchProduct(req)
		h.log.Debug("Searchs encontrados: ", modeloSearchs)
		return c.JSON(http.StatusOK, modeloSearchs)
	}
}

func (h *ProductController) GetRecomendationsProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info("Fetching recommendations for user")
		user := c.Get("user").(jwt.MapClaims)
		userId := user["sub"].(string)
		h.log.Debug("User ID: ", userId)
		items := h.productService.GetRecomendationsProduct(c.Request().Context(), userId)
		return c.JSON(http.StatusOK, items)
	}
}

func (h *ProductController) GetCategories() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info("Fetching categories")
		categories, err := h.productService.GetCategories()
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

func (h *ProductController) CreatePosteoDemo() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info("Creando posteo")
		posteo := demo.PostDTO{}
		if err := c.Bind(&posteo); err != nil {
			h.log.Error("Error al parsear el body: ", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "no se pudo parsear el body",
			})
		}
		h.log.Debug("Posteo recibido: ", posteo)
		claims := c.Get("user").(jwt.MapClaims)
		sub := claims["sub"].(string)
		h.log.Debug("Usuario: ", sub)
		userId, err := uuid.Parse(sub)
		if err != nil {
			h.log.Error("Error al parsear el userId: ", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "userId inválido",
			})
		}
		h.log.Debug("Posteo recibido: ", posteo)
		err = h.postsService.CreatePosteo(&posteo, userId.String())
		if err != nil {
			h.log.Error("Error al crear el posteo: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "no se pudo crear el posteo",
			})
		}
		h.log.Info("Posteo creado exitosamente")
		return c.JSON(http.StatusCreated, "Posteo creado exitosamente")
	}
}

func (h *ProductController) GetPosteosDemo() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info("GetPosteos Entrando")
		h.log.Debug("Obteniendo posteos")
		productId := c.QueryParams().Get("productId")
		h.log.Debug("productId recibido: ", productId)
		posteos := h.postsService.GetPosteos(productId)
		h.log.Debug("Posteos encontrados: ", posteos)
		return c.JSON(http.StatusOK, posteos)
	}
}
