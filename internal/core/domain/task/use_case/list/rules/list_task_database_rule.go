package rules

import (
	"context"
	"sword-project/internal/models"
	"sword-project/internal/repositories"
)

type ListTasksDatabaseRule struct {
	ctx            context.Context
	taskRepository repositories.TaskRepository
	params         *models.ListTasksParams
}

func NewListTasksDatabaseRule(ctx context.Context, taskRepository repositories.TaskRepository, params *models.ListTasksParams) *ListTasksDatabaseRule {
	return &ListTasksDatabaseRule{
		ctx:            ctx,
		taskRepository: taskRepository,
		params:         params,
	}
}

func (c *ListTasksDatabaseRule) Apply() (*[]models.Task, error) {
	return c.taskRepository.ListTasks(c.ctx, c.params)
}
