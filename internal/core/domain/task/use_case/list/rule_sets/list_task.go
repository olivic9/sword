package rule_sets

import (
	"sword-project/internal/core/domain/task/use_case/list/rules"
	"sword-project/internal/models"
)

type ListTaskRuleSet struct {
	listTaskRule rules.ListTasksDatabaseRule
}

func NewListTaskRuleSet(listTasksRule rules.ListTasksDatabaseRule) *ListTaskRuleSet {
	return &ListTaskRuleSet{
		listTaskRule: listTasksRule,
	}
}

func (c *ListTaskRuleSet) Apply() (*[]models.Task, error) {
	return c.listTaskRule.Apply()
}
