package mutator

import (
	"iter"
	"maps"
	"slices"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
)

type Collection struct {
	mutators map[string]Appendable
}

func NewCollection(normalMutators, activeMutators []NamedUnit) Collection {
	normal := makeNormalMutatorMap(normalMutators)
	active := makeActiveMutatorMap(activeMutators)

	names := slc.Union(slices.Collect(maps.Keys(normal)), slices.Collect(maps.Keys(active)))

	mutators := make(map[string]Appendable)
	for _, name := range names {
		m, exists := normal[name]
		if !exists {
			m = NewAppendable(name, NewSwitch(nil), NewSwitch(nil))
		}

		mutators[name] = m.ConcatActive(active[name].ActiveSwitch())
	}

	rc := Collection{mutators: mutators}
	rc.selfLink()
	return rc
}

func (c Collection) GetMutator(name string) optional.Of[base.Mutator] {
	m, exists := c.mutators[name]
	return optional.New[base.Mutator](m, exists)
}

func (c Collection) LinkClass(class *base.Class) {
	if mutator, exists := c.mutators[class.Name()]; exists {
		class.Link(mutator)
	}
}

func (c Collection) AppendNormal(mutator NamedUnit) Appendable {
	name := mutator.Name()
	sw := NewSwitchFromSingleObjectMutator(mutator)
	m, exists := c.mutators[name]
	if !exists {
		c.mutators[name] = NewAppendable(name, sw, NewSwitch(nil))
	} else {
		c.mutators[name] = m.ConcatNormal(sw)
	}
	return c.mutators[name]
}

func (c Collection) IterMutators() iter.Seq2[string, Appendable] {
	return func(yield func(string, Appendable) bool) {
		for name, mutator := range c.mutators {
			if !yield(name, mutator) {
				return
			}
		}
	}
}

func (c Collection) LoadGlobal(builtin Collection) {
	for _, mutator := range c.mutators {
		mutator.LinkClass(builtin)
	}
	maps.Copy(c.mutators, builtin.mutators)
}

func (c Collection) selfLink() {
	for _, mutator := range c.mutators {
		mutator.LinkClass(c)
	}
}

func makeNormalMutatorMap(ms []NamedUnit) map[string]Appendable {
	switches := makeSwitchMapByName(ms)

	mutators := make(map[string]Appendable)
	for name, sw := range switches {
		mutators[name] = NewAppendable(name, sw, NewSwitch(nil))
	}
	return mutators
}

func makeActiveMutatorMap(ms []NamedUnit) map[string]Appendable {
	switches := makeSwitchMapByName(ms)

	mutators := make(map[string]Appendable)
	for name, sw := range switches {
		mutators[name] = NewAppendable(name, NewSwitch(nil), sw)
	}
	return mutators
}

func makeSwitchMapByName(ms []NamedUnit) map[string]Switch {
	mutators := make(map[string]Switch)
	for _, m := range ms {
		name := m.Name()
		mutators[name] = mutators[name].Append(m)
	}
	return mutators
}
