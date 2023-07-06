package app

import (
	"database/sql"
	"sword-project/internal/services"
	"sword-project/pkg/logging"
)

type Application struct {
	Services services.Services
}

func NewApplication(db *sql.DB) (*Application, func()) {

	appServices := services.NewServices(
		services.NewCreateTaskService(db),
		services.NewListTasksService(db),
		services.NewFinishTaskService(db))

	app := &Application{
		Services: appServices,
	}

	return app, func() {
		logging.Logger.Sync()
		_ = db.Close()
	}
}
