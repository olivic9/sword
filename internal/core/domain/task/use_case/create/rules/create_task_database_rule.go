package rules

import (
	"context"
	"sword-project/internal/models"
	"sword-project/internal/repositories"
)

type CreateTaskDatabaseRule struct {
	ctx            context.Context
	taskRepository repositories.TaskRepository
	model          *models.Task
}

func NewCreateTaskDatabaseRule(ctx context.Context, taskRepository repositories.TaskRepository, model *models.Task) *CreateTaskDatabaseRule {
	return &CreateTaskDatabaseRule{
		ctx:            ctx,
		taskRepository: taskRepository,
		model:          model,
	}
}

func (c *CreateTaskDatabaseRule) Apply() error {
	return c.taskRepository.NewTask(c.ctx, c.model)
}
