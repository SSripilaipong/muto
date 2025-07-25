package parameter

import (
	"maps"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/go-common/rods"

	"github.com/SSripilaipong/muto/core/base"
)

type Parameter struct {
	variableMapping     map[string]VariableMapping
	variadicVarMapping  map[string]VariadicVarMapping
	remainingParamChain base.ParamChain
}

func New() *Parameter {
	return &Parameter{
		variableMapping:     make(map[string]VariableMapping),
		variadicVarMapping:  make(map[string]VariadicVarMapping),
		remainingParamChain: base.NewParamChain(nil),
	}
}

func (m *Parameter) Merge(n *Parameter) optional.Of[*Parameter] {
	m1, ok := m.mergeVariableMappings(n.variableMapping).Return()
	if !ok {
		return optional.Empty[*Parameter]()
	}
	m2, ok := m1.mergeVariadicVarMappings(n.variadicVarMapping).Return()
	if !ok {
		return optional.Empty[*Parameter]()
	}
	return m2.mergeRemainingParamChain(n.remainingParamChain)
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

func (m *Parameter) WithVariableMapping(x VariableMapping) optional.Of[*Parameter] {
	newM := m.Clone()
	k := x.name
	if y, exists := newM.variableMapping[k]; !exists || y == x {
		newM.variableMapping[k] = x
		return optional.Value(newM)
	}
	return optional.Empty[*Parameter]()
}

func (m *Parameter) WithVariableMappings(xs []VariableMapping) optional.Of[*Parameter] {
	p := m.Clone()
	for _, vm := range xs {
		q, ok := p.WithVariableMapping(vm).Return()
		if !ok {
			return optional.Empty[*Parameter]()
		}
		p = q
	}
	return optional.Value(p)
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

func (m *Parameter) mergeRemainingParamChain(paramChain base.ParamChain) optional.Of[*Parameter] {
	newM := m.Clone()
	newM.remainingParamChain = newM.remainingParamChain.AppendAll(paramChain)
	return optional.Value(newM)
}

func (m *Parameter) WithVariadicVarMapping(x VariadicVarMapping) optional.Of[*Parameter] {
	newM := m.Clone()
	k := x.name
	if y, exists := newM.variadicVarMapping[k]; !exists || x.Equal(y) {
		newM.variadicVarMapping[k] = x
		return optional.Value(newM)
	}
	return optional.Empty[*Parameter]()
}

func (m *Parameter) WithVariadicVarMappings(xs []VariadicVarMapping) optional.Of[*Parameter] {
	p := m.Clone()
	for _, vm := range xs {
		q, ok := p.WithVariadicVarMapping(vm).Return()
		if !ok {
			return optional.Empty[*Parameter]()
		}
		p = q
	}
	return optional.Value(p)
}

func (m *Parameter) Clone() *Parameter {
	return &Parameter{
		variableMapping:     maps.Clone(m.variableMapping),
		variadicVarMapping:  maps.Clone(m.variadicVarMapping),
		remainingParamChain: m.remainingParamChain.Clone(),
	}
}

func (m *Parameter) VariableMappings() []VariableMapping {
	var result []VariableMapping
	for _, v := range m.variableMapping {
		result = append(result, v)
	}
	return result
}

func (m *Parameter) VariableMap() rods.Map[string, base.Node] {
	result := make(map[string]base.Node)
	for _, v := range m.variableMapping {
		result[v.Name()] = v.Node()
	}
	return rods.NewMap(result)
}

func (m *Parameter) VariadicVarMappings() []VariadicVarMapping {
	var result []VariadicVarMapping
	for _, v := range m.variadicVarMapping {
		result = append(result, v)
	}
	return result
}

func (m *Parameter) VariadicVarMap() rods.Map[string, []base.Node] {
	result := make(map[string][]base.Node)
	for _, v := range m.variadicVarMapping {
		result[v.Name()] = v.Nodes()
	}
	return rods.NewMap(result)
}

func (m *Parameter) SetRemainingParamChain(params base.ParamChain) *Parameter {
	newM := m.Clone()
	newM.remainingParamChain = params
	return newM
}

func (m *Parameter) AddRemainingParamChain(nodes []base.Node) *Parameter {
	newM := m.Clone()
	newM.remainingParamChain = newM.remainingParamChain.Append(nodes)
	return newM
}

func (m *Parameter) RemainingParamChain() base.ParamChain {
	return m.remainingParamChain
}

func (m *Parameter) AppendAllRemainingParamChain(remaining base.ParamChain) *Parameter {
	newM := m.Clone()
	newM.remainingParamChain = newM.remainingParamChain.AppendAll(remaining)
	return newM
}

func NewParameterWithVariableMapping(mapping VariableMapping) *Parameter {
	m := New().mergeVariableMappings(map[string]VariableMapping{mapping.name: mapping})
	if m.IsEmpty() {
		panic("wtf")
	}
	return m.Value()
}

func NewParameterWithVariadicVarMapping(mapping VariadicVarMapping) *Parameter {
	m := New().mergeVariadicVarMappings(map[string]VariadicVarMapping{mapping.name: mapping})
	if m.IsEmpty() {
		panic("wtf")
	}
	return m.Value()
}

func SetRemainingParamChain(params base.ParamChain) func(*Parameter) *Parameter {
	return func(parameter *Parameter) *Parameter {
		return parameter.SetRemainingParamChain(params)
	}
}

func AddRemainingParamChain(nodes []base.Node) func(*Parameter) *Parameter {
	return func(parameter *Parameter) *Parameter {
		return parameter.AddRemainingParamChain(nodes)
	}
}

func WithVariadicVarMappings(x VariadicVarMapping) func(parameter *Parameter) optional.Of[*Parameter] {
	return func(parameter *Parameter) optional.Of[*Parameter] {
		return parameter.WithVariadicVarMapping(x)
	}
}

func Merge(m, n *Parameter) optional.Of[*Parameter] {
	return m.Merge(n)
}
