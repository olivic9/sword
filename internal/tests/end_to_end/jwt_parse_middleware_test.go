package end_to_end

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sword-project/pkg/configs"
	"sword-project/pkg/server/http_server/http_middlewares"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func TestJwtParseMiddleware(t *testing.T) {

	configs.InitializeConfigs()
	configs.ApplicationCfg.JwtSecret = secretKey

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.Use(http_middlewares.JwtParseMiddleware())

	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	t.Run("No token", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusUnauthorized {
			t.Fatalf("Expected status %d, got %d", http.StatusUnauthorized, resp.Code)
		}
	})

	// Test invalid token
	t.Run("Invalid token", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "invalid_token")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusUnauthorized {
			t.Fatalf("Expected status %d, got %d", http.StatusUnauthorized, resp.Code)
		}
	})

	t.Run("Valid token", func(t *testing.T) {

		validToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		})

		token, _ := validToken.SignedString([]byte(secretKey))

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", token)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Fatalf("Expected status %d, got %d", http.StatusOK, resp.Code)
		}
	})

	t.Run("Expired token", func(t *testing.T) {

		expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
		})

		token, _ := expiredToken.SignedString([]byte(secretKey))

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", token)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusUnauthorized {
			t.Fatalf("Expected status 401 but got %d", resp.Code)
		}

		responseData, _ := io.ReadAll(resp.Body)
		jsonResponse := make(map[string]interface{})
		json.Unmarshal(responseData, &jsonResponse)

		if errorString, ok := jsonResponse["error"].(string); ok {
			if !strings.Contains(errorString, "token is expired") {
				t.Errorf("Expected error message to contain 'token is expired', got %s", errorString)
			}
		} else {
			t.Errorf("Expected error message, got %v", jsonResponse["error"])
		}
	})
}
