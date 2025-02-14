package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
)

type headVariableBuilder struct {
	name string
}

func newHeadVariableBuilder(variable st.Variable) headVariableBuilder {
	return headVariableBuilder{name: variable.Name()}
}

func (b headVariableBuilder) Build(mutation *parameter.Parameter) optional.Of[base.Node] {
	x, ok := mutation.VariableValue(b.name).Return()
	if !ok {
		return optional.Empty[base.Node]()
	}

	if base.IsObjectNode(x) {
		obj := base.UnsafeNodeToObject(x)
		if base.IsPrimitiveObject(obj) {
			return optional.Value(obj.Head())
		}
	}

	return optional.Value(x)
}
