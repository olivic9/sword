package rule_sets

import (
	"sword-project/internal/core/domain/task/use_case/get/rules"
	"sword-project/internal/models"
)

type GetTaskRuleSet struct {
	getTaskRule rules.GetTaskDatabaseRule
}

func NewGetTaskRuleSet(getTaskRule rules.GetTaskDatabaseRule) *GetTaskRuleSet {
	return &GetTaskRuleSet{
		getTaskRule: getTaskRule,
	}
}

func (c *GetTaskRuleSet) Apply() (*models.Task, error) {
	return c.getTaskRule.Apply()
}
