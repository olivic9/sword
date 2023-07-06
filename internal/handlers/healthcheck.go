package handlers

import (
	"sword-project/pkg/configs"
	"sword-project/pkg/responses"

	"github.com/gin-gonic/gin"
)

func HealthcheckHandler(c *gin.Context) {
	responses.OkResponse(c, gin.H{
		"status": "available",
		"system_info": map[string]string{
			"environment": configs.ApplicationCfg.Env,
		},
	})
}
