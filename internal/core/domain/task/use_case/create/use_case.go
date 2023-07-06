package create

import (
	"context"
	"sword-project/internal/core/domain/task/use_case/create/rule_sets"
	"sword-project/internal/core/domain/task/use_case/create/rules"
	"sword-project/internal/models"
	"sword-project/internal/repositories"
)

type UseCase struct {
	taskRepository repositories.TaskRepository
}

func NewCreateTaskUseCase(taskRepository repositories.TaskRepository) *UseCase {
	return &UseCase{
		taskRepository: taskRepository,
	}
}

func (u *UseCase) Execute(ctx context.Context, model *models.Task) error {
	return rule_sets.NewCreateTaskRuleSet(
		*rules.NewCreateTaskDatabaseRule(ctx, u.taskRepository, model),
	).Apply()
}
