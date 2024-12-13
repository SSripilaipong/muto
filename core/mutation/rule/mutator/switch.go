package mutator

import (
	"slices"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type Switch struct {
	mutators []ObjectMutator
}

func NewSwitch(mutators []ObjectMutator) Switch {
	return Switch{mutators: slices.Clone(mutators)}
}

func (s Switch) Mutate(obj base.Object) optional.Of[base.Node] {
	for _, mutator := range s.mutators {
		if result := mutator.Mutate(obj); result.IsNotEmpty() {
			return result
		}
	}
	return optional.Empty[base.Node]()
}

func (s Switch) Append(mutator ObjectMutator) Switch {
	s.mutators = append(s.mutators, mutator)
	return s
}

func (s Switch) Mutators() []ObjectMutator {
	return s.mutators
}

func (s Switch) SetGlobalMutator(gm NameBasedMutator) {
	for _, m := range s.mutators {
		if r, isGlobalAware := m.(GlobalMutatorAware); isGlobalAware {
			r.SetGlobalMutator(gm)
		}
	}
}

var _ ObjectMutator = Switch{}
