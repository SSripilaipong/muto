package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
)

type variableBuilder struct {
	name string
}

func newVariableBuilder(variable st.Variable) variableBuilder {
	return variableBuilder{name: variable.Name()}
}

func (b variableBuilder) Build(mutation *parameter.Parameter) optional.Of[base.Node] {
	return mutation.VariableValue(b.name)
}
