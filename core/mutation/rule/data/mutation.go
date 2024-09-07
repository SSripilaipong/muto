package data

import (
	"maps"
	"slices"

	"muto/common/optional"
	"muto/core/base"
)

type Mutation struct {
	variableMapping   map[string]VariableMapping
	remainingChildren []base.Node
}

func NewMutation() *Mutation {
	return &Mutation{
		variableMapping: make(map[string]VariableMapping),
	}
}

func (m *Mutation) Merge(n *Mutation) optional.Of[*Mutation] {
	return m.mergeVariableMappings(n.variableMapping)
}

func (m *Mutation) VariableValue(name string) optional.Of[base.Node] {
	variable, exists := m.variableMapping[name]
	return optional.New(variable.node, exists)
}

func (m *Mutation) mergeVariableMappings(mapping map[string]VariableMapping) optional.Of[*Mutation] {
	newM := m.Clone()
	for k, x := range mapping {
		if y, ok := newM.variableMapping[k]; !ok {
			newM.variableMapping[k] = x
		} else if x != y {
			return optional.Empty[*Mutation]()
		}
	}
	return optional.Value(newM)
}

func (m *Mutation) Clone() *Mutation {
	return &Mutation{
		variableMapping:   maps.Clone(m.variableMapping),
		remainingChildren: slices.Clone(m.remainingChildren),
	}
}

func (m *Mutation) WithRemainingChildren(nodes []base.Node) *Mutation {
	newM := m.Clone()
	if len(nodes) > 0 {
		newM.remainingChildren = nodes
	}
	return newM
}

func (m *Mutation) RemainingChildren() []base.Node {
	return m.remainingChildren
}

func NewMutationWithVariableMapping(mapping VariableMapping) *Mutation {
	m := NewMutation().mergeVariableMappings(map[string]VariableMapping{mapping.name: mapping})
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
