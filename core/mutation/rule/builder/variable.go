package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

type paramVariableBuilder struct {
	name string
}

func newParamVariableBuilder(variable st.Variable) paramVariableBuilder {
	return paramVariableBuilder{name: variable.Name()}
}

func (b paramVariableBuilder) Build(mutation *parameter.Parameter) optional.Of[base.Node] {
	return mutation.VariableValue(b.name)
}

func (b paramVariableBuilder) DisplayString() string {
	return b.name
}
