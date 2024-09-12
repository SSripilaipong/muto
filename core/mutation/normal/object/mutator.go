package object

import (
	"muto/common/optional"
	"muto/core/base"
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

func (t Mutator) Mutate(obj base.NamedObject) optional.Of[base.Node] {
	if t.name != obj.Name() {
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
