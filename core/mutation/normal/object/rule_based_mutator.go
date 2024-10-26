package object

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type RuleBasedMutator struct {
	name          string
	mutationRules []func(t base.Object) optional.Of[base.Node]
}

func NewRuleBasedMutator(name string, mutationRules []func(t base.Object) optional.Of[base.Node]) RuleBasedMutator {
	return RuleBasedMutator{
		name:          name,
		mutationRules: mutationRules,
	}
}

func MergeRuleBasedMutators(mutators ...RuleBasedMutator) RuleBasedMutator {
	name := mutators[0].name
	var rules []func(t base.Object) optional.Of[base.Node]
	for _, mutator := range mutators {
		if mutator.Name() != name {
			panic("mutator name mismatched")
		}
		rules = append(rules, mutator.mutationRules...)
	}
	return NewRuleBasedMutator(name, rules)
}

func (t RuleBasedMutator) Mutate(name string, obj base.Object) optional.Of[base.Node] {
	if name != t.name {
		return optional.Empty[base.Node]()
	}
	for _, mutate := range t.mutationRules {
		if result := mutate(obj); result.IsNotEmpty() {
			return result
		}
	}
	return optional.Empty[base.Node]()
}

func (t RuleBasedMutator) Name() string {
	return t.name
}

func (t RuleBasedMutator) SetGlobalMutator(_ Mutator) {}
