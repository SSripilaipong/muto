package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/data"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func buildBoolean(x st.Boolean) func(*data.Mutation) optional.Of[base.Node] {
	value := base.NewBoolean(x.BooleanValue())

	return func(mapping *data.Mutation) optional.Of[base.Node] {
		return optional.Value[base.Node](value)
	}
}
