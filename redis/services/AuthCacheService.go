package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tcero76/marketplace/bff-service/services"
	logger "github.com/tcero76/marketplace/config/log"
	"github.com/tcero76/marketplace/redis/model"
)

type AuthCacheService struct {
	Rdb *redis.Client
	log *logger.LoggerLogstash
}

func NewAuthCacheService(log *logger.LoggerLogstash, rdb *redis.Client) services.IAuthCacheService {
	log.Info("Initializing AuthCacheService")
	return &AuthCacheService{Rdb: rdb, log: log}
}

func (h *AuthCacheService) SaveSession(key string, s model.SessionData, ctx context.Context) error {
	h.log.FromContext(ctx).Info("Saving session to Redis")
	h.log.FromContext(ctx).Info("SaveSession arg Key =", key)
	if key == "" {
		h.log.FromContext(ctx).Error("Error: key is empty")
		return nil
	}
	botSessionJSON, err := json.Marshal(s.BotSession)
	if err != nil {
		return errors.New("marshal bot session: " + err.Error())
	}
	data := map[string]interface{}{
		"access_token":     s.AccessToken,
		"refresh_token":    s.RefreshToken,
		"user_id":          s.UserID,
		"idp":              s.Idp,
		"session_id":       s.SessionID,
		"is_authenticated": s.IsAuthenticated,
		"login_challenge":  s.LoginChallenge,
		"bot_session":      string(botSessionJSON),
	}

	err = h.Rdb.HSet(ctx, "session:"+key, data).Err()
	if err != nil {
		h.log.FromContext(ctx).Error("Error saving session to Redis: ", err)
	}
	return err
}

func (h *AuthCacheService) GetSession(key string, ctx context.Context) (*model.SessionData, error) {
	h.log.FromContext(ctx).Info("Getting session from Redis")
	h.log.FromContext(ctx).Info("GetSession arg Key =", key)
	m, err := h.Rdb.HGetAll(context.Background(), "session:"+key).Result()
	if err != nil {
		h.log.FromContext(ctx).Error("Error getting session from Redis: ", err)
		return nil, err
	}
	if len(m) == 0 {
		return nil, errors.New("no session found for key: " + key)
	}
	var botSession *model.FormSession
	if raw := m["bot_session"]; raw != "" {
		botSession = &model.FormSession{}

		if err := json.Unmarshal([]byte(raw), botSession); err != nil {
			return nil, fmt.Errorf("decoding bot_session: %w", err)
		}
	}
	s := &model.SessionData{
		AccessToken:     m["access_token"],
		RefreshToken:    m["refresh_token"],
		UserID:          m["user_id"],
		Idp:             m["idp"],
		SessionID:       m["session_id"],
		LoginChallenge:  m["login_challenge"],
		IsAuthenticated: m["is_authenticated"] == "true",
		BotSession:      botSession,
	}
	return s, nil
}

func (h *AuthCacheService) StoreTokenInRedis(sessionID string, key string, s *model.FormSession, ctx context.Context) error {
	h.log.Info("Storing in Redis")
	h.log.Debug("sessionID=", sessionID, " key=", key)
	botSessionJSON, err := json.Marshal(s)
	if err != nil {
		return errors.New("marshal bot session: " + err.Error())
	}
	err = h.Rdb.HSet(
		ctx,
		"session:"+sessionID,
		key, botSessionJSON,
	).Err()
	if err != nil {
		h.log.Error("Error storing in Redis: ", err)
		return err
	}
	err = h.Rdb.Expire(ctx,
		"session:"+sessionID,
		720*time.Hour).Err()
	if err != nil {
		h.log.Error("Error setting expiration in Redis: ", err)
		return err
	}
	return nil
}

func (h *AuthCacheService) LoadTokenFromRedis(sessionID string, key string, ctx context.Context) (string, error) {
	h.log.Info("Loading from Redis: sessionID=", sessionID, " key=", key)
	data, err := h.Rdb.HGet(ctx,
		"session:"+sessionID,
		key).Result()
	if err != nil {
		h.log.Error("Error loading from Redis: ", err)
		return "", err
	}
	return data, nil
}

func (h *AuthCacheService) LoadSessionAll(sessionID string, ctx context.Context) (map[string]string, error) {
	h.log.Info("Loading all from Redis: sessionID=", sessionID)
	data, err := h.Rdb.HGetAll(ctx, "session:"+sessionID).Result()
	if err != nil {
		h.log.Error("Error loading all from Redis: ", err)
		return nil, err
	}
	return data, nil
}

func (h *AuthCacheService) SaveSessionAll(sessionID string, session map[string]string, ctx context.Context) error {
	h.log.Info("Saving all to Redis: sessionID=", sessionID)
	err := h.Rdb.HSet(ctx,
		"session:"+sessionID,
		session).Err()
	if err != nil {
		h.log.Error("Error saving all to Redis: ", err)
	}
	return err
}
