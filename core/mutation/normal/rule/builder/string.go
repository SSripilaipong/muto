package builder

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/mutation/normal/rule/data"
	st "muto/syntaxtree"
)

func buildString(s st.String) func(mapping *data.Mutation) optional.Of[base.Node] {
	value := s.Value()
	return func(mapping *data.Mutation) optional.Of[base.Node] {
		return optional.Value[base.Node](base.NewString(value))
	}
}
