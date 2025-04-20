package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type ParamVariable struct {
	name string
}

func NewParamVariable(name string) ParamVariable {
	return ParamVariable{name: name}
}

func (v ParamVariable) Extract(x base.Node) optional.Of[*parameter.Parameter] {
	return optional.Value(parameter.NewParameterWithVariableMapping(parameter.NewVariableMapping(v.Name(), x)))
}

func (v ParamVariable) Name() string {
	return v.name
}

var _ NodeExtractor = ParamVariable{}

type IgnoredParamVariable struct {
	name string
}

func NewIgnoredParamVariable() IgnoredParamVariable {
	return IgnoredParamVariable{}
}

func (v IgnoredParamVariable) Extract(base.Node) optional.Of[*parameter.Parameter] {
	return optional.Value(parameter.New())
}

var _ NodeExtractor = IgnoredParamVariable{}
