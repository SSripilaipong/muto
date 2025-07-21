package extractor

import (
	"fmt"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type Rune struct {
	value rune
}

func NewRune(value rune) Rune {
	return Rune{value: value}
}

func (s Rune) Extract(node base.Node) optional.Of[*parameter.Parameter] {
	if base.IsRuneNode(node) && base.UnsafeNodeToRune(node).Value() == s.Value() {
		return optional.Value(parameter.New())
	}
	return optional.Empty[*parameter.Parameter]()
}

func (s Rune) Value() rune {
	return s.value
}

func (s Rune) DisplayString() string {
	return fmt.Sprintf("%#v", s.Value())
}

var _ NodeExtractor = Rune{}
