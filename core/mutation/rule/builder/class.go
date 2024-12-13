package builder

import (
	"github.com/SSripilaipong/muto/core/base"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
)

func newClassBuilder(x st.Class) constantBuilder[base.Class] {
	return newConstantBuilder(base.NewClass(x.Name()))
}
