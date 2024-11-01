package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func buildTag(x st.Tag) func(*parameter.Parameter) optional.Of[base.Node] {
	value := base.NewTag(x.Name())

	return func(mapping *parameter.Parameter) optional.Of[base.Node] {
		return optional.Value[base.Node](value)
	}
}
