package extractor

import (
	"fmt"

	"github.com/SSripilaipong/go-common/optional"

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

func (s String) DisplayString() string {
	return fmt.Sprintf("%#v", s.Value())
}

var _ NodeExtractor = String{}
