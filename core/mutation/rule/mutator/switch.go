package mutator

import (
	"slices"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
)

type Switch struct {
	mutators []ObjectMutator
}

func NewSwitch(mutators []ObjectMutator) Switch {
	return Switch{mutators: slices.Clone(mutators)}
}

func NewSwitchFromSingleObjectMutator(mutator ObjectMutator) Switch {
	return NewSwitch(slc.Pure[ObjectMutator](mutator))
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
	return NewSwitch(slc.CloneAppend(s.mutators, mutator))
}

func (s Switch) Mutators() []ObjectMutator {
	return s.mutators
}

func (s Switch) Concat(t Switch) Switch {
	return NewSwitch(slc.CloneConcat(s.mutators, t.mutators))
}

var _ ObjectMutator = Switch{}
