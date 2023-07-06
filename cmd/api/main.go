package main

import (
	"context"
	"sword-project/internal/app"
	"sword-project/pkg/configs"
	"sword-project/pkg/databases"
	"sword-project/pkg/logging"
	"sword-project/pkg/server/http_server"
)

var ctx = context.Background()

func main() {
	configs.InitializeConfigs()

	logging.InitializeApplicationLogger()
	apiApp, flush := app.NewApplication(databases.GetMysqlDatabase(ctx))
	defer flush()
	http_server.RunHttpServer(apiApp)
}
