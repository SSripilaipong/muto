package command

import (
	ruleMutator "github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

type AddRuleCommand struct {
	TypeMixin
	mutator ruleMutator.NamedObjectMutator
}

func (c AddRuleCommand) Rule() ruleMutator.NamedObjectMutator {
	return c.mutator
}

func NewAddRule(mutator ruleMutator.NamedObjectMutator) *AddRuleCommand {
	return &AddRuleCommand{
		TypeMixin: NewTypeMixin(TypeAddRule),
		mutator:   mutator,
	}
}

func UnsafeCommandToAddRuleCommand(cmd Command) *AddRuleCommand {
	return cmd.(*AddRuleCommand)
}
