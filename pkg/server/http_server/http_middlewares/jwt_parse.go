package http_middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sword-project/internal/models"
	"sword-project/pkg/configs"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtParseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		authToken := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)

		token, err := jwt.ParseWithClaims(authToken, &models.Token{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("unexpected signing method: %v", token.Header["alg"])})
			}

			return []byte(configs.ApplicationCfg.JwtSecret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		ctx = context.WithValue(ctx, "user", token.Claims.(*models.Token))

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
