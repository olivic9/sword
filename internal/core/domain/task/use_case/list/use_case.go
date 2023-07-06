package list

import (
	"context"
	"sword-project/internal/core/domain/task/use_case/list/rule_sets"
	"sword-project/internal/core/domain/task/use_case/list/rules"
	"sword-project/internal/models"
	"sword-project/internal/repositories"
)

type UseCase struct {
	taskRepository repositories.TaskRepository
}

func NewListTasksUseCase(taskRepository repositories.TaskRepository) *UseCase {
	return &UseCase{
		taskRepository: taskRepository,
	}
}

func (u *UseCase) Execute(ctx context.Context, params *models.ListTasksParams) (*[]models.Task, error) {
	return rule_sets.NewListTaskRuleSet(
		*rules.NewListTasksDatabaseRule(ctx, u.taskRepository, params),
	).Apply()
}
