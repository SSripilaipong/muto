package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func buildBoolean(x st.Boolean) func(*parameter.Parameter) optional.Of[base.Node] {
	value := base.NewBoolean(x.BooleanValue())

	return func(mapping *parameter.Parameter) optional.Of[base.Node] {
		return optional.Value[base.Node](value)
	}
}
