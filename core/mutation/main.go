package mutation

import (
	"muto/common/optional"
	"muto/core/base"
	activeMutation "muto/core/mutation/active"
	normalMutation "muto/core/mutation/normal"
	st "muto/syntaxtree"
)

func NewFromStatements(ss []st.Statement) func(base.Object) optional.Of[base.Node] {
	active := activeMutation.NewFromStatements(ss)
	normal := normalMutation.NewFromStatements(ss)
	return func(obj base.Object) optional.Of[base.Node] {
		if result, ok := active(obj).Return(); ok {
			return optional.Value(result)
		}
		return normal(obj)
	}
}
