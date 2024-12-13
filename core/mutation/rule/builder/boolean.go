package builder

import (
	"github.com/SSripilaipong/muto/core/base"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
)

func newBooleanBuilder(x st.Boolean) constantBuilder[base.Boolean] {
	return newConstantBuilder(base.NewBoolean(x.BooleanValue()))
}
