package builder

import (
	"github.com/SSripilaipong/muto/core/base"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
)

func newTagBuilder(x st.Tag) constantBuilder[base.Node] {
	return newConstantBuilder(base.NewTag(x.Name()))
}
