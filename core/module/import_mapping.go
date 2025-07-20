package module

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/rods"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/portal"
)

type ImportMapping struct {
	mapping map[string]Module
}

func NewImportMapping(mapping rods.Map[string, Module]) ImportMapping {
	return ImportMapping{mapping: mapping.ToMap()}
}

func (m ImportMapping) GetCollection(name string) optional.Of[mutator.Collection] {
	mod, exists := m.mapping[name]
	if !exists {
		return optional.Empty[mutator.Collection]()
	}
	return optional.Value(mod.MutatorCollection())
}

func (m ImportMapping) LoadGlobal(global Module) {
	for _, mod := range m.mapping {
		mod.LoadGlobal(global)
	}
}

func (m ImportMapping) MountPortal(q *portal.Portal) {
	for _, mod := range m.mapping {
		mod.MountPortal(q)
	}
}
