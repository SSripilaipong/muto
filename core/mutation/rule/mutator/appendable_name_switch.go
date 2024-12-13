package mutator

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type AppendableNameSwitch struct {
	mutators map[string]Switch
}

func NewAppendableNameSwitch(ms []NamedObjectMutator) AppendableNameSwitch {
	mutators := make(map[string]Switch)
	for _, m := range ms {
		name := m.Name()
		mutators[name] = mutators[name].Append(m)
	}

	m := AppendableNameSwitch{mutators: mutators}
	m.SetGlobalMutator(m)
	return m
}

func (s AppendableNameSwitch) MutateByName(name string, obj base.Object) optional.Of[base.Node] {
	if m, ok := s.mutators[name]; ok {
		return m.Mutate(obj)
	}
	return optional.Empty[base.Node]()
}

func (s AppendableNameSwitch) Append(r NamedObjectMutator) {
	name := r.Name()
	m := s.mutators[name]
	s.mutators[name] = m.Append(r)
}

func (s AppendableNameSwitch) SetGlobalMutator(gm NameBasedMutator) {
	for _, m := range s.mutators {
		m.SetGlobalMutator(gm)
	}
}
