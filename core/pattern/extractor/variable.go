package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type VariableFactory struct {
}

func NewVariableFactory() VariableFactory {
	return VariableFactory{}
}

func (m VariableFactory) FixedVariable(name string) NodeExtractor {
	return NewVariable(name)
}

func (m VariableFactory) VariadicVariable(name string) NodeListExtractor {
	return NewContextFreeVariadic(name)
}

type Variable struct {
	name string
}

func NewVariable(name string) Variable {
	return Variable{name: name}
}

func (v Variable) Extract(x base.Node) optional.Of[*parameter.Parameter] {
	return optional.Value(parameter.NewParameterWithVariableMapping(parameter.NewVariableMapping(v.Name(), x)))
}

func (v Variable) Name() string {
	return v.name
}

var _ NodeExtractor = Variable{}

type ContextFreeVariadic struct {
	name string
}

func NewContextFreeVariadic(name string) ContextFreeVariadic {
	return ContextFreeVariadic{name: name}
}

func (v ContextFreeVariadic) Extract(nodes []base.Node) optional.Of[*parameter.Parameter] {
	return optional.Value(parameter.NewParameterWithVariadicVarMapping(parameter.NewVariadicVarMapping(v.Name(), nodes)))
}

func (v ContextFreeVariadic) Name() string {
	return v.name
}

var _ NodeListExtractor = ContextFreeVariadic{}

type ContextFreeIgnoreVariadic struct {
}

func NewContextFreeIgnoreVariadic() ContextFreeIgnoreVariadic {
	return ContextFreeIgnoreVariadic{}
}

func (v ContextFreeIgnoreVariadic) Extract([]base.Node) optional.Of[*parameter.Parameter] {
	return optional.Value(parameter.New())
}

var _ NodeListExtractor = ContextFreeVariadic{}
