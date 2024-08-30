package builder

import (
	"phi-lang/common/optional"
	"phi-lang/core/base"
	"phi-lang/core/mutation/rule/data"
	st "phi-lang/syntaxtree"
)

func buildString(s st.String) func(mapping *data.Mutation) optional.Of[base.Node] {
	value := s.Value()
	return func(mapping *data.Mutation) optional.Of[base.Node] {
		return optional.Value[base.Node](base.NewString(value))
	}
}
