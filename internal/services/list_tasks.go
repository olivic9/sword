package services

import (
	"context"
	"database/sql"
	"sword-project/internal/core/domain/task/use_case/list"
	"sword-project/internal/models"
	"sword-project/internal/repositories"
)

type ListTasksService struct {
	MysqlDb *sql.DB
}

func NewListTasksService(db *sql.DB) *ListTasksService {
	return &ListTasksService{
		MysqlDb: db,
	}
}

func (c *ListTasksService) List(ctx context.Context, params *models.ListTasksParams) (*[]models.Task, error) {

	return list.NewListTasksUseCase(*repositories.NewTaskRepository(c.MysqlDb)).Execute(ctx, params)
}
