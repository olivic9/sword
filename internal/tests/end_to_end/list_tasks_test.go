package end_to_end

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sword-project/internal/app"
	"sword-project/internal/handlers"
	"sword-project/pkg/configs"
	"sword-project/pkg/helpers/test_helpers"
	"sword-project/pkg/server/http_server/http_middlewares"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
)

func TestListTasksForManager(t *testing.T) {

	configs.InitializeConfigs()
	configs.ApplicationCfg.JwtSecret = jwtSecret
	asOf, _ := time.Parse("2006-01-02 15:04:05", "2023-01-01 00:00:00")
	expected := `{"data":[{"ID":1,"Title":"test","Summary":"test","Status":"pending","TeamID":0,"AssignedTechnicianID":0, "AssignedTechnicianName":"", "CreatedAt":"2023-01-01T00:00:00Z","FinishedAt":"2023-01-01T00:00:00Z"}]}`
	var expectedObject interface{}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	apiApp, flush := app.NewApplication(db)
	defer flush()

	rows := sqlmock.NewRows([]string{"id", "title", "summary", "status", "created_at", "finished_at"}).AddRow(1, "test", "test", "pending", asOf, asOf)
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT id, title, summary, status, created_at, finished_at FROM tasks WHERE team_id = ? LIMIT ? OFFSET ?`,
	)).WithArgs(1, 10, 1).WillReturnRows(rows)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.Use(http_middlewares.JwtParseMiddleware())

	router.GET("/api/tasks", handlers.NewApiHandler(apiApp.Services).ListTasks)

	t.Run("Happy Path", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/tasks", nil)
		req.Header.Set("Authorization", managerTeam1Token)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Fatalf("Expected status %d, got %d", http.StatusCreated, resp.Code)
		}

		responseData, _ := io.ReadAll(resp.Body)
		jsonResponse := make(map[string]interface{})
		json.Unmarshal(responseData, &jsonResponse)

		json.Unmarshal([]byte(expected), &expectedObject)

		test_helpers.IsObjectEqual(t, jsonResponse, expectedObject)
	})
}

func TestListTasksForTechnician(t *testing.T) {

	configs.InitializeConfigs()
	configs.ApplicationCfg.JwtSecret = jwtSecret
	asOf, _ := time.Parse("2006-01-02 15:04:05", "2023-01-01 00:00:00")
	expected := `{"data":[{"ID":1,"Title":"test","Summary":"test","Status":"pending","TeamID":0,"AssignedTechnicianID":0, "AssignedTechnicianName":"", "CreatedAt":"2023-01-01T00:00:00Z","FinishedAt":"2023-01-01T00:00:00Z"}]}`
	var expectedObject interface{}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	apiApp, flush := app.NewApplication(db)
	defer flush()

	rows := sqlmock.NewRows([]string{"id", "title", "summary", "status", "created_at", "finished_at"}).AddRow(1, "test", "test", "pending", asOf, asOf)
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT t.id,  t.title, t.summary, t.status, t.created_at, t.finished_at
					FROM tasks as t
					LEFT JOIN users ON t.assigned_technician_id = users.id
					WHERE t.assigned_technician_id IS NULL OR users.uuid = ?
					LIMIT ? 
					OFFSET ?`,
	)).WithArgs("c558a80a-f319-4c10-95d4-4282ef745b4b", 10, 1).WillReturnRows(rows)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.Use(http_middlewares.JwtParseMiddleware())

	router.GET("/api/tasks", handlers.NewApiHandler(apiApp.Services).ListTasks)

	t.Run("Happy Path", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/tasks", nil)
		req.Header.Set("Authorization", technician1Team1Token)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Fatalf("Expected status %d, got %d", http.StatusCreated, resp.Code)
		}

		responseData, _ := io.ReadAll(resp.Body)
		jsonResponse := make(map[string]interface{})
		json.Unmarshal(responseData, &jsonResponse)

		json.Unmarshal([]byte(expected), &expectedObject)

		test_helpers.IsObjectEqual(t, jsonResponse, expectedObject)
	})
}
