package extractor

import (
	"fmt"

	"github.com/SSripilaipong/go-common/optional"

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

func (v Variable) DisplayString() string {
	return v.Name()
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

func (v ContextFreeVariadic) DisplayString() string {
	return fmt.Sprintf("%s...", v.Name())
}

var _ NodeListExtractor = ContextFreeVariadic{}

type ContextFreeIgnoreVariadic struct {
	name string
}

func NewContextFreeIgnoreVariadic(name string) ContextFreeIgnoreVariadic {
	return ContextFreeIgnoreVariadic{name: name}
}

func (v ContextFreeIgnoreVariadic) Extract([]base.Node) optional.Of[*parameter.Parameter] {
	return optional.Value(parameter.New())
}

func (v ContextFreeIgnoreVariadic) DisplayString() string {
	return fmt.Sprintf("%s...", v.name)
}

var _ NodeListExtractor = ContextFreeVariadic{}
