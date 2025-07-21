package mutator

import (
	"slices"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/portal"
)

type Switch struct {
	mutators []Unit
}

func NewSwitch(mutators []Unit) Switch {
	return Switch{mutators: slices.Clone(mutators)}
}

func NewSwitchFromSingleObjectMutator(mutator Unit) Switch {
	return NewSwitch(slc.Pure[Unit](mutator))
}

func (s Switch) Mutate(obj base.Object) optional.Of[base.Node] {
	for _, mutator := range s.mutators {
		if result := mutator.Mutate(obj); result.IsNotEmpty() {
			return result
		}
	}
	return optional.Empty[base.Node]()
}

func (s Switch) Append(mutator Unit) Switch {
	return NewSwitch(slc.CloneAppend(s.mutators, mutator))
}

func (s Switch) Mutators() []Unit {
	return s.mutators
}

func (s Switch) Concat(t Switch) Switch {
	return NewSwitch(slc.CloneConcat(s.mutators, t.mutators))
}

func (s Switch) VisitClass(linker ClassVisitor) {
	for _, mutator := range s.mutators {
		mutator.VisitClass(linker)
	}
}

func (s Switch) MountPortal(p *portal.Portal) {
	for _, mutator := range s.mutators {
		portal.MountPortalToMutator(p, mutator)
	}
}

func MergeSwitches(switches ...Switch) Switch {
	var result []Unit
	for _, sw := range switches {
		result = append(result, sw.Mutators()...)
	}
	return NewSwitch(result)
}

var _ Unit = Switch{}

type NamedSwitch struct {
	Switch
	name string
}

func NewNamedSwitch(name string, sw Switch) NamedSwitch {
	return NamedSwitch{Switch: sw, name: name}
}

func (t NamedSwitch) Name() string {
	return t.name
}

func MergeNamedSwitches(mutators ...NamedSwitch) NamedSwitch {
	name := mutators[0].name
	var switches []Switch
	for _, mutator := range mutators {
		if mutator.Name() != name {
			panic("mutator name mismatched")
		}
		switches = append(switches, mutator.Switch)
	}
	return NewNamedSwitch(name, MergeSwitches(switches...))
}
