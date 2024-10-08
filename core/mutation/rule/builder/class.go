package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/data"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func buildClass(x st.Class) func(*data.Mutation) optional.Of[base.Node] {
	value := base.NewClass(x.Name())

	return func(r *data.Mutation) optional.Of[base.Node] {
		return optional.Value[base.Node](value)
	}
}
