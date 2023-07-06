package http_server

import (
	"net/http"
	"sword-project/pkg/configs"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	trustedOrigins := configs.ApplicationCfg.CorsTrustedOrigins

	return cors.New(cors.Config{
		AllowOrigins:     trustedOrigins,
		AllowMethods:     []string{http.MethodOptions, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Cache-Control", "X-XSRF-TOKEN"},
		AllowCredentials: true,
		AllowWildcard:    true,
		MaxAge:           60 * time.Second,
	})
}
