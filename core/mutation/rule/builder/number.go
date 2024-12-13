package builder

import (
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/base/datatype"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
)

func newNumberBuilder(x st.Number) constantBuilder[base.Node] {
	return newConstantBuilder(base.NewNumber(datatype.NewNumber(x.Value())))
}
