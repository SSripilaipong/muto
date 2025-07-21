package extractor

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type IgnoredParamVariable struct {
	name string
}

func NewIgnoredParamVariable(name string) IgnoredParamVariable {
	return IgnoredParamVariable{name: name}
}

func (v IgnoredParamVariable) Extract(base.Node) optional.Of[*parameter.Parameter] {
	return optional.Value(parameter.New())
}

func (v IgnoredParamVariable) DisplayString() string {
	return v.name
}

var _ NodeExtractor = IgnoredParamVariable{}
