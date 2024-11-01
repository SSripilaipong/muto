package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/base/datatype"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type Number struct {
	value datatype.Number
}

func NewNumber(value datatype.Number) Number {
	return Number{value: value}
}

func (n Number) Extract(node base.Node) optional.Of[*parameter.Parameter] {
	if base.IsNumberNode(node) && base.UnsafeNodeToNumber(node).Value() == n.Value() {
		return optional.Value(parameter.New())
	}
	return optional.Empty[*parameter.Parameter]()
}

func (n Number) Value() datatype.Number {
	return n.value
}

var _ NodeExtractor = Number{}
