package extractor

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type Boolean struct {
	value bool
}

func (b Boolean) Extract(x base.Node) optional.Of[*parameter.Parameter] {
	if base.IsBooleanNode(x) && base.UnsafeNodeToBoolean(x).Value() == b.Value() {
		return optional.Value(parameter.New())
	}
	return optional.Empty[*parameter.Parameter]()
}

func (b Boolean) Value() bool {
	return b.value
}

func (b Boolean) DisplayString() string {
	return fmt.Sprint(b.Value())
}

func NewBoolean(value bool) Boolean {
	return Boolean{value: value}
}

var _ NodeExtractor = Boolean{}
