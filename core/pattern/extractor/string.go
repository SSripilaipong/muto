package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type String struct {
	value string
}

func NewString(value string) String {
	return String{value: value}
}

func (s String) Extract(node base.Node) optional.Of[*parameter.Parameter] {
	if base.IsStringNode(node) && base.UnsafeNodeToString(node).Value() == s.Value() {
		return optional.Value(parameter.New())
	}
	return optional.Empty[*parameter.Parameter]()
}

func (s String) Value() string {
	return s.value
}

var _ NodeExtractor = String{}
