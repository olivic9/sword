package test_helpers

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sword-project/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func TestJwtParseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		token := models.Token{
			GivenName:        "Test",
			Surname:          "Subject",
			Email:            "t.t@test.com",
			Role:             "test",
			Team:             "1",
			UUID:             "aaa-aaa-aa-aa-a",
			RegisteredClaims: jwt.RegisteredClaims{},
		}

		//nolint
		ctx = context.WithValue(ctx, "user", &token)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func BuildGinTestEngine(w *httptest.ResponseRecorder) (*gin.Engine, *gin.Context) {

	gin.SetMode(gin.TestMode)
	e := gin.Default()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return e, ctx
}

func MockJsonGet(c *gin.Context, params gin.Params, u url.Values) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")

	c.Params = params

	c.Request.URL.RawQuery = u.Encode()
}

func MockJsonPost(c *gin.Context, jsonString string) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	var json = []byte(jsonString)

	c.Request.Body = io.NopCloser(bytes.NewBuffer(json))
}
