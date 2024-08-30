package builder

import (
	"phi-lang/common/optional"
	"phi-lang/core/base"
	"phi-lang/core/mutation/rule/data"
	st "phi-lang/syntaxtree"
)

func buildVariable(variable st.Variable) func(*data.Mutation) optional.Of[base.Node] {
	name := variable.Name()

	return func(mutation *data.Mutation) optional.Of[base.Node] {
		return mutation.VariableValue(name)
	}
}
