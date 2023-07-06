package rules

import (
	"context"
	"sword-project/internal/adapters"
	"sword-project/internal/models"
	"sword-project/pkg/configs"
	"time"
)

type PublishCreatedNegotiationOfferRule struct {
	kafkaAdapter adapters.KafkaService
	ctx          context.Context
	params       *models.FinishTaskParams
}

func NewPublishCreatedNegotiationOfferRule(kafkaAdapter adapters.KafkaService, ctx context.Context, params *models.FinishTaskParams) *PublishCreatedNegotiationOfferRule {
	return &PublishCreatedNegotiationOfferRule{
		kafkaAdapter: kafkaAdapter,
		ctx:          ctx,
		params:       params,
	}
}

func (p *PublishCreatedNegotiationOfferRule) Apply() {
	p.kafkaAdapter.ProduceMessage(models.CreateFinishedTaskKafkaMessage(&models.FinishedTask{
		ID:         p.params.TaskID,
		FinishedAt: time.Now(),
	}), p.ctx, configs.KafkaCfg.NotificationTopic)
}
