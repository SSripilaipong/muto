package mutator

import (
	"iter"
	"maps"
	"slices"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
)

type RuleCollection struct {
	mutators map[string]AppendableNamedRuleMutator
}

func NewRuleCollection(normalMutators, activeMutators []NamedObjectMutator) RuleCollection {
	normal := makeNormalRuleMutatorMap(normalMutators)
	active := makeActiveRuleMutatorMap(activeMutators)

	names := slc.Union(slices.Collect(maps.Keys(normal)), slices.Collect(maps.Keys(active)))

	mutators := make(map[string]AppendableNamedRuleMutator)
	for _, name := range names {
		m, exists := normal[name]
		if !exists {
			m = NewAppendableNamedRuleMutator(name, NewSwitch(nil), NewSwitch(nil))
		}

		mutators[name] = m.ConcatActive(active[name].ActiveSwitch())
	}

	return RuleCollection{mutators: mutators}
}

func (c RuleCollection) Active(name string, obj base.Object) optional.Of[base.Node] {
	m, exists := c.mutators[name]
	if !exists {
		return optional.Empty[base.Node]()
	}
	return m.Active(obj)
}

func (c RuleCollection) Normal(name string, obj base.Object) optional.Of[base.Node] {
	m, exists := c.mutators[name]
	if !exists {
		return optional.Empty[base.Node]()
	}
	return m.Normal(obj)
}

func (c RuleCollection) AppendNormalRule(mutator NamedObjectMutator) AppendableNamedRuleMutator {
	name := mutator.Name()
	sw := NewSwitchFromSingleObjectMutator(mutator)
	m, exists := c.mutators[name]
	if !exists {
		c.mutators[name] = NewAppendableNamedRuleMutator(name, sw, NewSwitch(nil))
	} else {
		c.mutators[name] = m.ConcatNormal(sw)
	}
	return c.mutators[name]
}

func (c RuleCollection) IterMutators() iter.Seq2[string, AppendableNamedRuleMutator] {
	return func(yield func(string, AppendableNamedRuleMutator) bool) {
		for name, mutator := range c.mutators {
			if !yield(name, mutator) {
				return
			}
		}
	}
}

func makeNormalRuleMutatorMap(ms []NamedObjectMutator) map[string]AppendableNamedRuleMutator {
	switches := makeSwitchMapByName(ms)

	mutators := make(map[string]AppendableNamedRuleMutator)
	for name, sw := range switches {
		mutators[name] = NewAppendableNamedRuleMutator(name, sw, NewSwitch(nil))
	}
	return mutators
}

func makeActiveRuleMutatorMap(ms []NamedObjectMutator) map[string]AppendableNamedRuleMutator {
	switches := makeSwitchMapByName(ms)

	mutators := make(map[string]AppendableNamedRuleMutator)
	for name, sw := range switches {
		mutators[name] = NewAppendableNamedRuleMutator(name, NewSwitch(nil), sw)
	}
	return mutators
}

func makeSwitchMapByName(ms []NamedObjectMutator) map[string]Switch {
	mutators := make(map[string]Switch)
	for _, m := range ms {
		name := m.Name()
		mutators[name] = mutators[name].Append(m)
	}
	return mutators
}
