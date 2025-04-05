package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type HeadVariable struct {
	name string
}

func NewHeadVariable(name string) HeadVariable {
	return HeadVariable{name: name}
}

func (v HeadVariable) Extract(x base.Node) optional.Of[*parameter.Parameter] {
	return optional.Value(parameter.NewParameterWithVariableMapping(parameter.NewVariableMapping(v.Name(), x)))
}

func (v HeadVariable) Name() string {
	return v.name
}

var _ NodeExtractor = HeadVariable{}
