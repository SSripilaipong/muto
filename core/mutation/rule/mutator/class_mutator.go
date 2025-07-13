package mutator

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type ClassMutator struct {
	name      string
	mutations Switch
}

func NewClassMutator(name string, rules []ObjectMutator) ClassMutator {
	return ClassMutator{
		name:      name,
		mutations: NewSwitch(rules),
	}
}

func MergeClassMutators(mutators ...ClassMutator) ClassMutator {
	name := mutators[0].name
	var rules []ObjectMutator
	for _, mutator := range mutators {
		if mutator.Name() != name {
			panic("mutator name mismatched")
		}
		rules = append(rules, mutator.mutations.Mutators()...)
	}
	return NewClassMutator(name, rules)
}

func (t ClassMutator) MutateByName(name string, obj base.Object) optional.Of[base.Node] {
	if name != t.name {
		return optional.Empty[base.Node]()
	}
	return t.Mutate(obj)
}

func (t ClassMutator) Mutate(obj base.Object) optional.Of[base.Node] {
	return t.mutations.Mutate(obj)
}

func (t ClassMutator) Append(mutator ObjectMutator) ClassMutator {
	t.mutations = t.mutations.Append(mutator)
	return t
}

func (t ClassMutator) LinkClass(linker ClassLinker) {
	t.mutations.LinkClass(linker)
}

func (t ClassMutator) Name() string {
	return t.name
}

func (t ClassMutator) SetName(name string) ClassMutator {
	t.name = name
	return t
}
