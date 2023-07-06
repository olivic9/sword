package get

import (
	"context"
	"sword-project/internal/core/domain/task/use_case/get/rule_sets"
	"sword-project/internal/core/domain/task/use_case/get/rules"
	"sword-project/internal/models"
	"sword-project/internal/repositories"
)

type UseCase struct {
	taskRepository *repositories.TaskRepository
}

func NewGetTaskUseCase(taskRepository *repositories.TaskRepository) *UseCase {
	return &UseCase{
		taskRepository: taskRepository,
	}
}

func (u *UseCase) Get(ctx context.Context, id int64) (*models.Task, error) {
	return rule_sets.NewGetTaskRuleSet(
		*rules.NewGetTaskDatabaseRule(ctx, u.taskRepository, id),
	).Apply()
}
