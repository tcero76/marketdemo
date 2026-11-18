package controller

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tcero76/marketplace/bff-service/dto"
	"github.com/tcero76/marketplace/bff-service/services"
	logger "github.com/tcero76/marketplace/config/log"
)

type PostController struct {
	log              *logger.LoggerLogstash
	inventoryService services.IInventory
	postsService     services.IDemoPosts
}

func NewPostController(log *logger.LoggerLogstash,
	inventoryService services.IInventory,
	postsService services.IDemoPosts) *PostController {
	return &PostController{log, inventoryService, postsService}
}

// CreatePosteoDemo godoc
//
// @Summary      Crear un posteo
// @Description  Crea un posteo asociado al usuario autenticado.
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        posteo body demo.PostDTO true "Datos del posteo"
// @Success      201 {string} string "Posteo creado exitosamente"
// @Failure      400 {object} map[string]string "Error en los datos enviados"
// @Failure      401 {object} map[string]string "No autorizado"
// @Failure      500 {object} map[string]string "Error al crear el posteo"
// @Security     BearerAuth
// @Router       /usuario/createPostDemo [post]
func (h *PostController) CreatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info("Creando posteo")
		posteo := dto.PostRequest{}
		if err := c.Bind(&posteo); err != nil {
			h.log.Error("Error al parsear el body: ", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "no se pudo parsear el body",
			})
		}
		h.log.Debug("Posteo recibido: ", posteo)
		claims, ok := c.Get("user").(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing user claims",
			})
		}
		h.log.Debug("claims: ", claims)
		sub, ok := claims["sub"].(string)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing subject claims",
			})
		}
		h.log.Debug("Subject: ", sub)
		userId, err := uuid.Parse(sub)
		if err != nil {
			h.log.Error("Error al parsear el userId: ", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "userId inválido",
			})
		}
		err = h.postsService.CreatePosteo(&posteo, userId.String())
		if err != nil {
			h.log.Error("Error al crear el posteo: ", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		h.log.Info("Posteo creado exitosamente")
		return c.JSON(http.StatusCreated, map[string]string{
			"status": "Posteo creado exitosamente",
		})
	}
}

// GetPosteosDemo godoc
//
// @Summary      Obtener posteos de un producto
// @Description  Obtiene los posteos asociados a un producto.
// @Tags         Posts
// @Produce      json
// @Param        productId query string true "Identificador del producto"
// @Success      200 {array} demo.PostDTO
// @Failure      400 {object} map[string]string "productId requerido"
// @Failure      500 {object} map[string]string "Error al obtener los posteos"
// @Security     BearerAuth
// @Router       /usuario/getPosteosDemo [get]
func (h *PostController) GetPosts() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info("GetPosteos Entrando")
		productId := c.QueryParams().Get("productId")
		if productId == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "productId requerido",
			})
		}
		h.log.Debug("productId recibido: ", productId)
		posteos, err := h.postsService.GetPosteos(productId)
		if err != nil {
			h.log.Error("Error al obtener los posteos: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "no se pudo obtener los posteos",
			})
		}
		h.log.Debug("Posteos encontrados: ", posteos)
		return c.JSON(http.StatusOK, posteos)
	}
}
