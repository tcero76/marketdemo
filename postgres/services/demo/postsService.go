package demo

import (
	"errors"

	"github.com/google/uuid"
	"github.com/tcero76/marketplace/bff-service/dto/demo"
	logger "github.com/tcero76/marketplace/config"
	"github.com/tcero76/marketplace/postgres/adapters"
	"github.com/tcero76/marketplace/postgres/model"
	"gorm.io/gorm"
)

type PostsService struct {
	dbWrite *gorm.DB
	dbRead  *gorm.DB
	log     *logger.LoggerLogstash
}

func NewPostsService(log *logger.LoggerLogstash, dbWrite *gorm.DB, dbRead *gorm.DB) *PostsService {
	return &PostsService{dbWrite: dbWrite, dbRead: dbRead, log: log}
}

func (c *PostsService) GetPosts(limit int, offset int) []demo.PostDTO {
	c.log.Info("GetPosts Entrando al servicio")
	var posts []model.PostsDemo
	err := c.dbRead.Limit(limit).
		Offset(offset).
		Find(&posts)
	if err.Error != nil {
		c.log.Error("Error al obtener los posts: ", err.Error)
		return []demo.PostDTO{}
	}
	postsDTO := adapters.ToPostDTOS(posts)
	return postsDTO
}

func (c *PostsService) GetTotalPosts() int64 {
	c.log.Info("GetTotalPosts Entrando al servicio")
	var total int64
	err := c.dbRead.Model(&model.Posts{}).
		Count(&total)
	if err.Error != nil {
		c.log.Error("Error al obtener el total de posts: ", err.Error)
		return 0
	}
	return total
}

func (c *PostsService) CreatePosteo(posteoDTO *demo.PostDTO, userId string) error {
	c.log.Info("CreatePosteo Entrando al servicio")
	posteo := adapters.ToPosteoModel(posteoDTO)
	c.log.Debug("Posteo a crear: ", posteo)
	id, err := uuid.Parse(userId)
	if err != nil {
		c.log.Error("Error al parsear el userId: ", err)
		return err
	}
	posteo.UserID = id
	if posteo.UserID.String() != userId {
		err := errors.New("el userId del posteo no coincide con el userId del token")
		c.log.Error(err)
		return err
	}
	result := c.dbWrite.Save(&posteo)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}

func (c *PostsService) GetPosteos(productId string) []demo.PostDTO {
	c.log.Info("GetPosteos Entrando al servicio modelo: ", productId)
	posteos := []model.PostsDemo{}
	err := c.dbRead.Where("product_id = ?", productId).Find(&posteos)
	if err.Error != nil {
		c.log.Error("Error al obtener los posteos: ", err.Error)
		return []demo.PostDTO{}
	}
	return adapters.ToPostDTOS(posteos)
}
