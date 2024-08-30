package base

import (
	"fmt"

	"phi-lang/core/base/datatype"
)

type Number struct {
	value datatype.Number
}

func NewNumber(value datatype.Number) Node {
	return Number{value: value}
}

func (n Number) NodeType() NodeType {
	return NodeTypeNumber
}

func (n Number) Value() datatype.Number {
	return n.value
}

func (n Number) String() string {
	if n.value.IsInt() {
		return fmt.Sprintf("%d", n.value.ToInt())
	}
	return fmt.Sprintf("%.2f", n.value.ToFloat())
}

func UnsafeNodeToNumber(n Node) Number {
	return n.(Number)
}
