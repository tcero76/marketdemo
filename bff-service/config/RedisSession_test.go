package config

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"

	"github.com/labstack/echo/v4"
	logConfig "github.com/tcero76/marketplace/config/log"
	"github.com/tcero76/marketplace/redis/model"
)

type mockAuthCacheService struct {
    GetSessionFunc  func(string, context.Context) (*model.SessionData, error)
    SaveSessionFunc func(string, model.SessionData, context.Context) error
	LoadTokenFromRedisFunc func(sessionID string, key string, ctx context.Context) (string, error)
}

func (m *mockAuthCacheService) GetSession(id string, ctx context.Context) (*model.SessionData, error) {
    return m.GetSessionFunc(id, ctx)
}

func (m *mockAuthCacheService) SaveSession(id string, s model.SessionData, ctx context.Context) error {
    return m.SaveSessionFunc(id, s, ctx)
}

func (m *mockAuthCacheService) LoadTokenFromRedis(sessionID string, key string, ctx context.Context) (string, error) {
	return m.LoadTokenFromRedisFunc(sessionID, key, ctx)
}

func TestRedisSessionMiddleware_NewSession(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	saveCalled := false
	mock := &mockAuthCacheService{
		GetSessionFunc: func(id string, ctx context.Context) (*model.SessionData, error) {
			return nil, errors.New("not found")
		},
		SaveSessionFunc: func(id string, s model.SessionData, ctx context.Context) error {
			saveCalled = true

			if id == "" {
				t.Fatal("session id vacío")
			}

			if s.SessionID != id {
				t.Fatal("session id incorrecto")
			}

			return nil
		},
	}
	nextCalled := false
	next := func(c echo.Context) error {
		nextCalled = true
		v := c.Get("session_data")
		if v == nil {
			t.Fatal("session_data no existe")
		}
		session := v.(*model.SessionData)
		if session.IsAuthenticated {
			t.Fatal("la sesión debería comenzar no autenticada")
		}
		return c.NoContent(http.StatusOK)
	}
	handler := RedisSessionMiddleware(
		mock,
		logConfig.NewLoggerLogstash("test"),
	)(next)
	if err := handler(c); err != nil {
		t.Fatal(err)
	}
	if !nextCalled {
		t.Fatal("next no fue ejecutado")
	}
	if !saveCalled {
		t.Fatal("SaveSession no fue llamado")
	}
	cookies := rec.Result().Cookies()
	if len(cookies) == 0 {
		t.Fatal("debía crear una cookie")
	}
	if cookies[0].Name != "session_id" {
		t.Fatal("cookie incorrecta")
	}
}

func TestRedisSessionMiddleware_ExistingSession(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "session_id",
		Value: "abc123",
	})
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	saveCalled := false
	mock := &mockAuthCacheService{
		GetSessionFunc: func(id string, ctx context.Context) (*model.SessionData, error) {

			return &model.SessionData{
				SessionID:       id,
				IsAuthenticated: true,
			}, nil
		},

		SaveSessionFunc: func(id string, s model.SessionData, ctx context.Context) error {
			saveCalled = true
			return nil
		},
	}
	next := func(c echo.Context) error {
		session := c.Get("session_data").(*model.SessionData)

		if !session.IsAuthenticated {
			t.Fatal("la sesión debería venir autenticada")
		}

		return c.NoContent(http.StatusOK)
	}
	handler := RedisSessionMiddleware(
		mock,
		logConfig.NewLoggerLogstash("test"),
	)(next)
	if err := handler(c); err != nil {
		t.Fatal(err)
	}
	if saveCalled {
		t.Fatal("SaveSession no debería ejecutarse")
	}
	if len(rec.Result().Cookies()) != 0 {
		t.Fatal("no debería crear cookies")
	}
}

func TestRedisSessionMiddleware_SessionNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "session_id",
		Value: "abc123",
	})
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	saveCalled := false
	mock := &mockAuthCacheService{
		GetSessionFunc: func(id string, ctx context.Context) (*model.SessionData, error) {
			return nil, errors.New("not found")
		},
		SaveSessionFunc: func(id string, s model.SessionData, ctx context.Context) error {
			saveCalled = true
			return nil
		},
	}
	next := func(c echo.Context) error {
		session := c.Get("session_data").(*model.SessionData)
		if session.SessionID != "abc123" {
			t.Fatal("session id incorrecto")
		}
		if session.IsAuthenticated {
			t.Fatal("debería iniciar como false")
		}
		return c.NoContent(http.StatusOK)
	}
	handler := RedisSessionMiddleware(
		mock,
		logConfig.NewLoggerLogstash("test"),
	)(next)
	if err := handler(c); err != nil {
		t.Fatal(err)
	}
	if !saveCalled {
		t.Fatal("SaveSession debería ejecutarse")
	}
	if len(rec.Result().Cookies()) != 0 {
		t.Fatal("no debería crear una nueva cookie")
	}
}

func TestRedisSessionMiddleware_SaveSessionError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	nextCalled := false
	mock := &mockAuthCacheService{
		GetSessionFunc: func(id string, ctx context.Context) (*model.SessionData, error) {
			return nil, errors.New("not found")
		},
		SaveSessionFunc: func(id string, s model.SessionData, ctx context.Context) error {
			return errors.New("redis down")
		},
	}
	next := func(c echo.Context) error {
		nextCalled = true
		return c.NoContent(http.StatusOK)
	}
	handler := RedisSessionMiddleware(
		mock,
		logConfig.NewLoggerLogstash("test"),
	)(next)
	err := handler(c)
	if err != nil {
		t.Fatal(err)
	}
	if nextCalled {
		t.Fatal("next no debería ejecutarse")
	}
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d, got %d",
			http.StatusInternalServerError,
			rec.Code)
	}
	expected :=  `{"error":"unable to create session"}`
	if strings.TrimSpace(rec.Body.String()) != expected {
		t.Fatalf("unexpected body: %s", rec.Body.String())
	}
	if c.Get("session_data") != nil {
		t.Fatal("session_data no debería existir")
	}
}