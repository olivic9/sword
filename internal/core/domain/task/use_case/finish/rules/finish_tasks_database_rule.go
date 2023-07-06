package rules

import (
	"context"
	"sword-project/internal/models"
	"sword-project/internal/repositories"
)

type FinishTaskDatabaseRule struct {
	ctx            context.Context
	taskRepository repositories.TaskRepository
	params         *models.FinishTaskParams
}

func NewFinishTaskDatabaseRule(ctx context.Context, taskRepository repositories.TaskRepository, params *models.FinishTaskParams) *FinishTaskDatabaseRule {
	return &FinishTaskDatabaseRule{
		ctx:            ctx,
		taskRepository: taskRepository,
		params:         params,
	}
}

func (c *FinishTaskDatabaseRule) Apply() error {
	return c.taskRepository.FinishTask(c.ctx, c.params)
}
