package builder

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/base/datatype"
	"muto/core/mutation/normal/rule/data"
	st "muto/syntaxtree"
)

func buildNumber(x st.Number) func(*data.Mutation) optional.Of[base.Node] {
	value := base.NewNumber(datatype.NewNumber(x.Value()))

	return func(mapping *data.Mutation) optional.Of[base.Node] {
		return optional.Value[base.Node](value)
	}
}
