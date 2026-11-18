package integration

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tcero76/marketplace/bff-service/server"
	logConfig "github.com/tcero76/marketplace/config/log"
	"github.com/tcero76/marketplace/postgres/config"
	postsModel "github.com/tcero76/marketplace/postgres/model/posts"
	"gorm.io/gorm"
)

func TestGetPosts(t *testing.T) {
	e, err := server.StartServer()
	require.NoError(t, err)
	t.Run("Test GetPosts Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet,
			"/usuario/getPosteos?productId=831ad86a-6956-475f-8a61-5a89d72b130d",
			nil)
		req.Header.Set("Authorization", "Bearer cualquier-cosa")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestCreatePost(t *testing.T) {
	log := logConfig.NewLoggerLogstash("🗄️ BFFTesting")
	db, err := config.GetPostgresWrite(log)
	assert.NoError(t, err)
	e, err := server.StartServer()
	require.NoError(t, err)
	t.Run("Test CreatePost Success", func(t *testing.T) {
		postID := uuid.New()
		body := `{
			"id":"` + postID.String() + `",
			"texto":"Lorem ipsum",
			"ProductId":"88a7c616-21ff-45bf-aa22-dbc6e3b0acb4"
		}`
		req := httptest.NewRequest(http.MethodPost,
			"/usuario/createPost",
			strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer cualquier-cosa")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		require.Equal(t, http.StatusCreated, rec.Code)
		deletePost(t, db, postID)
	})
}

func deletePost(t *testing.T, db *gorm.DB, postId uuid.UUID) error {
	t.Helper()
	if err := db.
		Where("id = ?", postId).
		Delete(&postsModel.Posts{}).Error; err != nil {
		return err
	}
	return nil
}
