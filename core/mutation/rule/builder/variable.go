package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/data"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func buildVariable(variable st.Variable) func(*data.Mutation) optional.Of[base.Node] {
	name := variable.Name()

	return func(mutation *data.Mutation) optional.Of[base.Node] {
		return mutation.VariableValue(name)
	}
}
