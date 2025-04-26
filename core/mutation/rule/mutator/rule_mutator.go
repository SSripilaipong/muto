package mutator

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type AppendableNamedRuleMutator struct {
	name   string
	active Switch
	normal Switch
}

func NewAppendableNamedRuleMutator(name string, normal, active Switch) AppendableNamedRuleMutator {
	return AppendableNamedRuleMutator{name: name, normal: normal, active: active}
}

func (m AppendableNamedRuleMutator) Active(obj base.Object) optional.Of[base.Node] {
	return m.active.Mutate(obj)
}

func (m AppendableNamedRuleMutator) Normal(obj base.Object) optional.Of[base.Node] {
	return m.normal.Mutate(obj)
}

func (m AppendableNamedRuleMutator) AppendNormal(mutator ObjectMutator) AppendableNamedRuleMutator {
	return NewAppendableNamedRuleMutator(m.name, m.normal.Append(mutator), m.active)
}

func (m AppendableNamedRuleMutator) AppendActive(mutator ObjectMutator) AppendableNamedRuleMutator {
	return NewAppendableNamedRuleMutator(m.name, m.normal, m.active.Append(mutator))
}

func (m AppendableNamedRuleMutator) ConcatNormal(sw Switch) AppendableNamedRuleMutator {
	return NewAppendableNamedRuleMutator(m.name, m.normal.Concat(sw), m.active)
}

func (m AppendableNamedRuleMutator) ConcatActive(sw Switch) AppendableNamedRuleMutator {
	return NewAppendableNamedRuleMutator(m.name, m.normal, m.active.Concat(sw))
}

func (m AppendableNamedRuleMutator) ActiveSwitch() Switch {
	return m.active
}

var _ RuleMutator = AppendableNamedRuleMutator{}
