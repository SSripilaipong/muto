package parameter

import (
	"maps"
	"slices"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type Parameter struct {
	variableMapping    map[string]VariableMapping
	variadicVarMapping map[string]VariadicVarMapping
	remainingChildren  []base.Node
}

func New() *Parameter {
	return &Parameter{
		variableMapping:    make(map[string]VariableMapping),
		variadicVarMapping: make(map[string]VariadicVarMapping),
	}
}

func (m *Parameter) Merge(n *Parameter) optional.Of[*Parameter] {
	m1, ok := m.mergeVariableMappings(n.variableMapping).Return()
	if !ok {
		return optional.Empty[*Parameter]()
	}
	return m1.mergeVariadicVarMappings(n.variadicVarMapping)
}

func (m *Parameter) VariableValue(name string) optional.Of[base.Node] {
	variable, exists := m.variableMapping[name]
	return optional.New(variable.node, exists)
}

func (m *Parameter) VariadicVarValue(name string) optional.Of[[]base.Node] {
	variable, exists := m.variadicVarMapping[name]
	return optional.New(variable.nodes, exists)
}

func (m *Parameter) mergeVariableMappings(mapping map[string]VariableMapping) optional.Of[*Parameter] {
	newM := m.Clone()
	for k, x := range mapping {
		if y, ok := newM.variableMapping[k]; !ok {
			newM.variableMapping[k] = x
		} else if !x.Equals(y) {
			return optional.Empty[*Parameter]()
		}
	}
	return optional.Value(newM)
}

func (m *Parameter) WithVariableMappings(x VariableMapping) optional.Of[*Parameter] {
	newM := m.Clone()
	k := x.name
	if y, exists := newM.variableMapping[k]; !exists || y == x {
		newM.variableMapping[k] = x
		return optional.Value(newM)
	}
	return optional.Empty[*Parameter]()
}

func (m *Parameter) mergeVariadicVarMappings(mapping map[string]VariadicVarMapping) optional.Of[*Parameter] {
	newM := m.Clone()
	for k, x := range mapping {
		if y, ok := newM.variadicVarMapping[k]; !ok {
			newM.variadicVarMapping[k] = x
		} else if !x.Equal(y) {
			return optional.Empty[*Parameter]()
		}
	}
	return optional.Value(newM)
}

func (m *Parameter) WithVariadicVarMappings(x VariadicVarMapping) optional.Of[*Parameter] {
	newM := m.Clone()
	k := x.name
	if y, exists := newM.variadicVarMapping[k]; !exists || x.Equal(y) {
		newM.variadicVarMapping[k] = x
		return optional.Value(newM)
	}
	return optional.Empty[*Parameter]()
}

func (m *Parameter) Clone() *Parameter {
	return &Parameter{
		variableMapping:    maps.Clone(m.variableMapping),
		variadicVarMapping: maps.Clone(m.variadicVarMapping),
		remainingChildren:  slices.Clone(m.remainingChildren),
	}
}

func (m *Parameter) WithRemainingChildren(nodes []base.Node) *Parameter {
	newM := m.Clone()
	if len(nodes) > 0 {
		newM.remainingChildren = nodes
	}
	return newM
}

func (m *Parameter) RemainingChildren() []base.Node {
	return m.remainingChildren
}

func NewParameterWithVariableMapping(mapping VariableMapping) *Parameter {
	m := New().mergeVariableMappings(map[string]VariableMapping{mapping.name: mapping})
	if m.IsEmpty() {
		panic("wtf")
	}
	return m.Value()
}

func WithRemainingChildren(nodes []base.Node) func(*Parameter) *Parameter {
	return func(parameter *Parameter) *Parameter {
		return parameter.WithRemainingChildren(nodes)
	}
}

func WithVariadicVarMappings(x VariadicVarMapping) func(parameter *Parameter) optional.Of[*Parameter] {
	return func(parameter *Parameter) optional.Of[*Parameter] {
		return parameter.WithVariadicVarMappings(x)
	}
}

func Merge(m, n *Parameter) optional.Of[*Parameter] {
	return m.Merge(n)
}
