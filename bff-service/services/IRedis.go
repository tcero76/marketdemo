package services

import (
	"context"

	"github.com/tcero76/marketplace/redis/model"
)

type IAuthCacheService interface {
	SaveSession(key string, s model.SessionData, ctx context.Context) error
	GetSession(key string, ctx context.Context) (*model.SessionData, error)
	StoreTokenInRedis(sessionID string, key string, s *model.FormSession, ctx context.Context) error
	LoadTokenFromRedis(sessionID string, key string, ctx context.Context) (string, error)
}
