package data

import (
	"maps"

	"phi-lang/common/optional"
	"phi-lang/core/base"
)

type Mutation struct {
	variableMapping map[string]VariableMapping
}

func (m Mutation) Merge(n Mutation) optional.Of[Mutation] {
	return m.mergeVariableMappings(n.variableMapping)
}

func (m Mutation) mergeVariableMappings(mapping map[string]VariableMapping) optional.Of[Mutation] {
	newMapping := maps.Clone(m.variableMapping)
	for k, x := range mapping {
		if y, ok := m.variableMapping[k]; !ok {
			newMapping[k] = x
		} else if x != y {
			return optional.Empty[Mutation]()
		}
	}
	m.variableMapping = newMapping
	return optional.Value(m)
}

func NewMutationWithVariableMapping(mapping VariableMapping) Mutation {
	m := Mutation{}.mergeVariableMappings(map[string]VariableMapping{mapping.name: mapping})
	if m.IsEmpty() {
		panic("wtf")
	}
	return m.Value()
}

type VariableMapping struct {
	name string
	node base.Node
}

func NewVariableMapping(name string, node base.Node) VariableMapping {
	return VariableMapping{
		name: name,
		node: node,
	}
}
