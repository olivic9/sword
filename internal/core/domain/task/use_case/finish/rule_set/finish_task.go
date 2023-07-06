package rule_sets

import (
	"sword-project/internal/core/domain/task/use_case/finish/rules"
)

type FinishTaskRuleSet struct {
	finishTaskDatabaseRule     rules.FinishTaskDatabaseRule
	publishFinishedTaskMessage rules.PublishCreatedNegotiationOfferRule
}

func NewFinishTaskRuleSet(finishTaskDatabaseRule rules.FinishTaskDatabaseRule,
	publishFinishedTaskMessage rules.PublishCreatedNegotiationOfferRule) *FinishTaskRuleSet {
	return &FinishTaskRuleSet{
		finishTaskDatabaseRule:     finishTaskDatabaseRule,
		publishFinishedTaskMessage: publishFinishedTaskMessage,
	}
}

func (c *FinishTaskRuleSet) Apply() error {
	err := c.finishTaskDatabaseRule.Apply()

	if err == nil {
		c.publishFinishedTaskMessage.Apply()
	}

	return err
}
