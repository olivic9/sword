package http_server

import (
	"context"
	"fmt"
	"net/http"
	"sword-project/internal/app"
	"sword-project/internal/routes"
	"sword-project/pkg/configs"
	"sword-project/pkg/logging"
	"time"

	"github.com/gin-gonic/gin"
)

func RunHttpServer(app *app.Application) {

	engine := gin.Default()

	srv := http.Server{
		Addr:              fmt.Sprintf(":%s", configs.ApplicationCfg.HttpServerPort),
		Handler:           routes.Routes(engine, app),
		ReadHeaderTimeout: 60 * time.Second,
	}

	logging.Logger.Info(fmt.Sprintf("Starting server on addr: %s", srv.Addr), logging.Metadata{})
	logging.Logger.Fatal(context.TODO(), srv.ListenAndServe(), logging.Metadata{})
}
