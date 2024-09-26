package object

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type Mutator struct {
	name          string
	mutationRules []func(t base.Object) optional.Of[base.Node]
}

func NewMutator(name string, mutationRules []func(t base.Object) optional.Of[base.Node]) Mutator {
	return Mutator{
		name:          name,
		mutationRules: mutationRules,
	}
}

func MergeMutators(mutators ...Mutator) Mutator {
	name := mutators[0].name
	var rules []func(t base.Object) optional.Of[base.Node]
	for _, mutator := range mutators {
		if mutator.Name() != name {
			panic("mutator name mismatched")
		}
		rules = append(rules, mutator.mutationRules...)
	}
	return NewMutator(name, rules)
}

func (t Mutator) Mutate(name string, obj base.Object) optional.Of[base.Node] {
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

func (t Mutator) Name() string {
	return t.name
}

func MutatorName(t Mutator) string {
	return t.Name()
}
