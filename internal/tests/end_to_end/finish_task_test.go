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

func TestFinishTask(t *testing.T) {
	configs.InitializeConfigs()
	configs.ApplicationCfg.JwtSecret = jwtSecret
	json := `{"task_id": 1}`

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	apiApp, flush := app.NewApplication(db)
	defer flush()

	mock.ExpectExec("UPDATE tasks SET").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.Use(http_middlewares.JwtParseMiddleware())

	router.PATCH("/api/task/finish", handlers.NewApiHandler(apiApp.Services).FinishTask)

	t.Run("Happy Path", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPatch, "/api/task/finish", bytes.NewBuffer([]byte(json)))
		req.Header.Set("Authorization", technician1Team1Token)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Fatalf("Expected status %d, got %d", http.StatusOK, resp.Code)
		}
	})

	t.Run("Without body", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPatch, "/api/task/finish", nil)
		req.Header.Set("Authorization", technician1Team1Token)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusBadRequest {
			t.Fatalf("Expected status %d, got %d", http.StatusBadRequest, resp.Code)
		}
		test_helpers.IsObjectEqual(t, resp.Body.Bytes(), []byte(`{"error":{"message":"invalid request"}}`))
	})

}

func TestFinishTaskWithHandlerError(t *testing.T) {
	configs.InitializeConfigs()
	configs.ApplicationCfg.JwtSecret = jwtSecret
	json := `{"task_id": 1}`

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	apiApp, flush := app.NewApplication(db)
	defer flush()

	mock.ExpectExec("UPDATE tasks SET").
		WithArgs(1).WillReturnError(errors.New("simulated error"))

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.Use(http_middlewares.JwtParseMiddleware())

	router.PATCH("/api/task/finish", handlers.NewApiHandler(apiApp.Services).FinishTask)

	t.Run("Database error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPatch, "/api/task/finish", bytes.NewBuffer([]byte(json)))
		req.Header.Set("Authorization", technician1Team1Token)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusInternalServerError {
			t.Fatalf("Expected status %d, got %d", http.StatusInternalServerError, resp.Code)
		}

		test_helpers.IsObjectEqual(t, resp.Body.Bytes(), []byte(`{"error":{"message":"internal error"}}`))
	})

}
