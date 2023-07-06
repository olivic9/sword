package rules

import (
	"context"
	"sword-project/internal/models"
	"sword-project/internal/repositories"
)

type GetTaskDatabaseRule struct {
	ctx            context.Context
	taskRepository *repositories.TaskRepository
	id             int64
}

func NewGetTaskDatabaseRule(ctx context.Context, taskRepository *repositories.TaskRepository, id int64) *GetTaskDatabaseRule {
	return &GetTaskDatabaseRule{
		ctx:            ctx,
		taskRepository: taskRepository,
		id:             id,
	}
}

func (c *GetTaskDatabaseRule) Apply() (*models.Task, error) {
	return c.taskRepository.GetTaskByID(c.ctx, c.id)
}
