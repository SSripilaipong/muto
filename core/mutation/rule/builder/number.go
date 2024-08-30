package builder

import (
	"phi-lang/common/optional"
	"phi-lang/core/base"
	"phi-lang/core/base/datatype"
	"phi-lang/core/mutation/rule/data"
	st "phi-lang/syntaxtree"
)

func buildNumber(x st.Number) func(*data.Mutation) optional.Of[base.Node] {
	value := base.NewNumber(datatype.NewNumber(x.Value()))

	return func(mapping *data.Mutation) optional.Of[base.Node] {
		return optional.Value[base.Node](value)
	}
}
