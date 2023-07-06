package end_to_end

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"sword-project/internal/app"
	"sword-project/internal/handlers"
	"sword-project/pkg/configs"
	"sword-project/pkg/helpers/test_helpers"
	"sword-project/pkg/server/http_server/http_middlewares"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
)

func TestCreateNewTask(t *testing.T) {
	configs.InitializeConfigs()
	configs.ApplicationCfg.JwtSecret = jwtSecret
	json := `{"title": "Test title","summary": "Test summary"}`

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	apiApp, flush := app.NewApplication(db)
	defer flush()

	mock.ExpectExec("INSERT INTO `tasks`").
		WithArgs("Test title", "Test summary", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.Use(http_middlewares.JwtParseMiddleware())

	router.POST("/api/task", handlers.NewApiHandler(apiApp.Services).NewTask)

	t.Run("Happy Path", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/api/task", bytes.NewBuffer([]byte(json)))
		req.Header.Set("Authorization", technician1Team1Token)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusCreated {
			t.Fatalf("Expected status %d, got %d", http.StatusCreated, resp.Code)
		}
	})

	t.Run("Without body", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/api/task", nil)
		req.Header.Set("Authorization", technician1Team1Token)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusBadRequest {
			t.Fatalf("Expected status %d, got %d", http.StatusBadRequest, resp.Code)
		}
		test_helpers.IsObjectEqual(t, resp.Body.Bytes(), []byte(`{"error":{"message":"invalid request"}}`))

	})

}

func TestCreateNewTaskWithHandlerError(t *testing.T) {
	configs.InitializeConfigs()
	configs.ApplicationCfg.JwtSecret = jwtSecret
	json := `{"title": "Test title","summary": "Test summary"}`

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	apiApp, flush := app.NewApplication(db)
	defer flush()

	mock.ExpectExec("INSERT INTO `tasks`").
		WithArgs("Test title", "Test summary", 1).WillReturnError(errors.New("simulated error"))

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.Use(http_middlewares.JwtParseMiddleware())

	router.POST("/api/task", handlers.NewApiHandler(apiApp.Services).NewTask)

	t.Run("Database error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/api/task", bytes.NewBuffer([]byte(json)))
		req.Header.Set("Authorization", technician1Team1Token)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusInternalServerError {
			t.Fatalf("Expected status %d, got %d", http.StatusInternalServerError, resp.Code)
		}

		test_helpers.IsObjectEqual(t, resp.Body.Bytes(), []byte(`{"error":{"message":"internal error"}}`))
	})

}
