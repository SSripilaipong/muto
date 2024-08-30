package object

import (
	"phi-lang/common/optional"
	"phi-lang/core/base"
)

type Mutator struct {
	name          string
	mutationRules []func(t base.ObjectLike) optional.Of[base.Node]
}

func NewMutator(name string, mutationRules []func(t base.ObjectLike) optional.Of[base.Node]) Mutator {
	return Mutator{
		name:          name,
		mutationRules: mutationRules,
	}
}

func (t Mutator) Mutate(obj base.ObjectLike) optional.Of[base.Node] {
	if t.name != obj.ClassName() {
		return optional.Empty[base.Node]()
	}
	if obj.IsTerminated() {
		return optional.Empty[base.Node]()
	}
	for _, mutate := range t.mutationRules {
		if result := mutate(obj); !result.IsEmpty() {
			node := result.Value()
			if base.IsObjectNode(node) && base.UnsafeNodeToObject(node).IsTerminated() {
				return optional.Empty[base.Node]()
			}
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
