package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func buildVariable(variable st.Variable) func(*parameter.Parameter) optional.Of[base.Node] {
	name := variable.Name()

	return func(mutation *parameter.Parameter) optional.Of[base.Node] {
		return mutation.VariableValue(name)
	}
}
