package services

import (
	"context"
	"database/sql"
	"sword-project/internal/core/domain/task/use_case/create"
	"sword-project/internal/models"
	"sword-project/internal/repositories"
)

type CreateTaskService struct {
	MysqlDb *sql.DB
}

func NewCreateTaskService(db *sql.DB) *CreateTaskService {
	return &CreateTaskService{
		MysqlDb: db,
	}
}

func (c *CreateTaskService) Execute(ctx context.Context, task *models.Task) error {

	return create.NewCreateTaskUseCase(*repositories.NewTaskRepository(c.MysqlDb)).Execute(ctx, task)
}
