package finish

import (
	"context"
	"sword-project/internal/adapters"
	rulesets "sword-project/internal/core/domain/task/use_case/finish/rule_set"
	"sword-project/internal/core/domain/task/use_case/finish/rules"
	"sword-project/internal/models"
	"sword-project/internal/repositories"
)

type UseCase struct {
	taskRepository repositories.TaskRepository
	kafkaAdapter   adapters.KafkaService
}

func NewFinishTaskUseCase(taskRepository repositories.TaskRepository, kafkaAdapter adapters.KafkaService) *UseCase {
	return &UseCase{
		taskRepository: taskRepository,
		kafkaAdapter:   kafkaAdapter,
	}
}

func (u *UseCase) Execute(ctx context.Context, params *models.FinishTaskParams) error {
	return rulesets.NewFinishTaskRuleSet(
		*rules.NewFinishTaskDatabaseRule(ctx, u.taskRepository, params),
		*rules.NewPublishCreatedNegotiationOfferRule(u.kafkaAdapter, ctx, params),
	).Apply()
}
