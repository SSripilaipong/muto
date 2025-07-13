package command

import (
	ruleMutator "github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

type AddRuleCommand struct {
	TypeMixin
	mutator ruleMutator.NamedUnit
}

func (c AddRuleCommand) Rule() ruleMutator.NamedUnit {
	return c.mutator
}

func NewAddRule(mutator ruleMutator.NamedUnit) *AddRuleCommand {
	return &AddRuleCommand{
		TypeMixin: NewTypeMixin(TypeAddRule),
		mutator:   mutator,
	}
}

func UnsafeCommandToAddRuleCommand(cmd Command) *AddRuleCommand {
	return cmd.(*AddRuleCommand)
}
