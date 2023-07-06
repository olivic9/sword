package routes

import (
	"sword-project/internal/app"
	"sword-project/internal/handlers"
	"sword-project/pkg/server/http_server/http_middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine, app *app.Application) *gin.Engine {

	engine.GET("/ping", handlers.HealthcheckHandler)

	rootPath := engine.Group("/")
	{
		api := rootPath.Group("api", http_middlewares.JwtParseMiddleware())
		{
			task := api.Group("task")
			{
				task.POST("/", handlers.NewApiHandler(app.Services).NewTask)
				task.PATCH("/finish", handlers.NewApiHandler(app.Services).FinishTask)
			}

			tasks := api.Group("tasks")
			{
				tasks.GET("/", handlers.NewApiHandler(app.Services).ListTasks)
			}

		}
	}

	return engine
}
