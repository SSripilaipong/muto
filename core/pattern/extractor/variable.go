package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

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
