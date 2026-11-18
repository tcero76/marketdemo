package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/tcero76/marketplace/bff-service/dto"
	"github.com/tcero76/marketplace/bff-service/mock"
	logConfig "github.com/tcero76/marketplace/config/log"
	"go.uber.org/mock/gomock"
)

func TestCreatePosteo(t *testing.T) {
	log := logConfig.NewLoggerLogstash("🗄️ BFFTesting")
	ctrl := gomock.NewController(t)
	mockPosts := mock.NewMockIDemoPosts(ctrl)
	controller := &PostController{
		log:          log,
		postsService: mockPosts,
	}
	e := echo.New()
	ID := uuid.New()
	reqDTO := dto.PostRequest{
		ID:    ID,
		Texto: "Lorem ipsum",
	}
	t.Run("Test CreatePosteo Success", func(t *testing.T) {
		expected := map[string]string{
			"status": "Posteo creado exitosamente",
		}
		body := `{
			"id":"` + ID.String() + `",
			"texto":"Lorem ipsum"
		}`
		mockPosts.EXPECT().
			CreatePosteo(gomock.Eq(&reqDTO), ID.String()).
			Return(nil).Times(1)
		req := httptest.NewRequest(
			http.MethodPost,
			"/createPost",
			strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("user", jwt.MapClaims{
			"sub": ID.String(),
		})
		controller.CreatePost()(c)
		assert.Equal(t, http.StatusCreated, rec.Code)
		var res map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		assert.NoError(t, err)
		assert.Equal(t, expected, res)
	})
	t.Run("Test CreatePosteo StatusUnauthorized", func(t *testing.T) {
		body := `{
			"id":"` + ID.String() + `",
			"texto":"Lorem ipsum"
		}`
		mockPosts.EXPECT().
			CreatePosteo(gomock.Eq(&reqDTO), ID.String()).
			Return(nil).Times(0)
		req := httptest.NewRequest(
			http.MethodPost,
			"/createPost",
			strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		controller.CreatePost()(c)
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		var res map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		assert.NoError(t, err)
		assert.Equal(t, "missing user claims", res["error"])
	})
}

func TestGetPosts(t *testing.T) {
	log := logConfig.NewLoggerLogstash("🗄️ BFFTesting")
	ctrl := gomock.NewController(t)
	mockPosts := mock.NewMockIDemoPosts(ctrl)
	controller := &PostController{
		log:          log,
		postsService: mockPosts,
	}
	e := echo.New()
	ID := uuid.New()
	// reqDTO := dto.PostRequest{
	// 	ID:    ID,
	// 	Texto: "Lorem ipsum",
	// }
	expected := []dto.PostResponse{
		{
			ID:    ID,
			Texto: "Lorem ipsum",
		},
	}
	t.Run("Test GetPosteos Success", func(t *testing.T) {
		body := `{
			"productId":"` + ID.String() + `"
		}`
		mockPosts.EXPECT().
			GetPosteos(ID.String()).
			Return(expected, nil).
			Times(1)
		req := httptest.NewRequest(
			http.MethodPost,
			"/getPosteos?productId="+ID.String(),
			strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		controller.GetPosts()(c)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
