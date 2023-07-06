package repositories

import (
	"database/sql"
)

type TaskRepository struct {
	database *sql.DB
}

func NewTaskRepository(database *sql.DB) *TaskRepository {
	return &TaskRepository{database: database}
}
