package base

import (
	"fmt"
	"strconv"

	"github.com/SSripilaipong/go-common/optional"
)

type Rune struct {
	value rune
}

func NewRune(value rune) Rune {
	return Rune{value: value}
}

func (Rune) NodeType() NodeType { return NodeTypeRune }

func (s Rune) MutateAsHead(params ParamChain) optional.Of[Node] {
	newParams := MutateParamChain(params)
	if newParams.IsNotEmpty() {
		return optional.Value[Node](NewCompoundObject(s, newParams.Value()))
	}
	return optional.Empty[Node]()
}

func (s Rune) Value() rune {
	return s.value
}

func (s Rune) TopLevelString() string {
	return s.String()
}

func (s Rune) String() string {
	return strconv.QuoteRune(s.value)
}

func (s Rune) MutoString() string {
	return fmt.Sprintf("%c", s.Value())
}

func (s Rune) Equals(t Rune) bool {
	return s.Value() == t.Value()
}

func UnsafeNodeToRune(n Node) Rune {
	return n.(Rune)
}

func RuneToNode(s Rune) Node {
	return s
}
