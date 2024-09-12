package builder

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/mutation/normal/rule/data"
	st "muto/syntaxtree"
)

func buildVariable(variable st.Variable) func(*data.Mutation) optional.Of[base.Node] {
	name := variable.Name()

	return func(mutation *data.Mutation) optional.Of[base.Node] {
		return mutation.VariableValue(name)
	}
}
