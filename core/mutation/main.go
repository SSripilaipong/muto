package mutation

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	activeMutation "github.com/SSripilaipong/muto/core/mutation/active"
	normalMutation "github.com/SSripilaipong/muto/core/mutation/normal"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func NewFromStatements(ss []st.Statement) func(base.MutableNode) optional.Of[base.Node] {
	mutate := mutation{
		active: activeMutation.NewFromStatements(ss),
		normal: normalMutation.NewFromStatements(ss),
	}

	return func(x base.MutableNode) optional.Of[base.Node] {
		return x.Mutate(mutate)
	}
}
