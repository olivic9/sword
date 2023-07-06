package rule_sets

import (
	"sword-project/internal/core/domain/task/use_case/create/rules"
)

type CreateTaskRuleSet struct {
	createTaskRule rules.CreateTaskDatabaseRule
}

func NewCreateTaskRuleSet(createTaskRule rules.CreateTaskDatabaseRule) *CreateTaskRuleSet {
	return &CreateTaskRuleSet{
		createTaskRule: createTaskRule,
	}
}

func (c *CreateTaskRuleSet) Apply() error {
	return c.createTaskRule.Apply()
}
