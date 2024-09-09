package data

import (
	"maps"
	"slices"

	"muto/common/optional"
	"muto/core/base"
)

type Mutation struct {
	variableMapping    map[string]VariableMapping
	variadicVarMapping map[string]VariadicVarMapping
	remainingChildren  []base.Node
}

func NewMutation() *Mutation {
	return &Mutation{
		variableMapping:    make(map[string]VariableMapping),
		variadicVarMapping: make(map[string]VariadicVarMapping),
	}
}

func (m *Mutation) Merge(n *Mutation) optional.Of[*Mutation] {
	m1, ok := m.mergeVariableMappings(n.variableMapping).Return()
	if !ok {
		return optional.Empty[*Mutation]()
	}
	return m1.mergeVariadicVarMappings(n.variadicVarMapping)
}

func (m *Mutation) VariableValue(name string) optional.Of[base.Node] {
	variable, exists := m.variableMapping[name]
	return optional.New(variable.node, exists)
}

func (m *Mutation) VariadicVarValue(name string) optional.Of[[]base.Node] {
	variable, exists := m.variadicVarMapping[name]
	return optional.New(variable.nodes, exists)
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

func (m *Mutation) WithVariableMappings(x VariableMapping) optional.Of[*Mutation] {
	newM := m.Clone()
	k := x.name
	if y, exists := newM.variableMapping[k]; !exists || y == x {
		newM.variableMapping[k] = x
		return optional.Value(newM)
	}
	return optional.Empty[*Mutation]()
}

func (m *Mutation) mergeVariadicVarMappings(mapping map[string]VariadicVarMapping) optional.Of[*Mutation] {
	newM := m.Clone()
	for k, x := range mapping {
		if y, ok := newM.variadicVarMapping[k]; !ok {
			newM.variadicVarMapping[k] = x
		} else if !x.Equal(y) {
			return optional.Empty[*Mutation]()
		}
	}
	return optional.Value(newM)
}

func (m *Mutation) WithVariadicVarMappings(x VariadicVarMapping) optional.Of[*Mutation] {
	newM := m.Clone()
	k := x.name
	if y, exists := newM.variadicVarMapping[k]; !exists || x.Equal(y) {
		newM.variadicVarMapping[k] = x
		return optional.Value(newM)
	}
	return optional.Empty[*Mutation]()
}

func (m *Mutation) Clone() *Mutation {
	return &Mutation{
		variableMapping:    maps.Clone(m.variableMapping),
		variadicVarMapping: maps.Clone(m.variadicVarMapping),
		remainingChildren:  slices.Clone(m.remainingChildren),
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

type VariadicVarMapping struct {
	name  string
	nodes []base.Node
}

func NewVariadicVarMapping(name string, nodes []base.Node) VariadicVarMapping {
	return VariadicVarMapping{
		name:  name,
		nodes: nodes,
	}
}

func (x VariadicVarMapping) Equal(y VariadicVarMapping) bool {
	if y.name != x.name {
		return false
	}
	if len(x.nodes) != len(y.nodes) {
		return false
	}
	for i := range x.nodes {
		if x.nodes[i] != y.nodes[i] {
			return false
		}
	}
	return true
}
