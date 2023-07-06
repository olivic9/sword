package services

import (
	"context"
	"database/sql"
	"sword-project/internal/adapters"
	"sword-project/internal/core/domain/task/use_case/finish"
	"sword-project/internal/models"
	"sword-project/internal/repositories"
	"sword-project/pkg/pubsub"
)

type FinishTaskService struct {
	MysqlDb *sql.DB
}

func NewFinishTaskService(db *sql.DB) *FinishTaskService {
	return &FinishTaskService{
		MysqlDb: db,
	}
}

func (c *FinishTaskService) Finish(ctx context.Context, params *models.FinishTaskParams) error {

	return finish.NewFinishTaskUseCase(*repositories.NewTaskRepository(c.MysqlDb),
		*adapters.NewKafkaService(pubsub.NewKafkaPublisher(ctx))).Execute(ctx, params)
}
