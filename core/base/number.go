package base

import (
	"fmt"

	"muto/core/base/datatype"
)

type Number struct {
	value datatype.Number
}

func NewNumber(value datatype.Number) Node {
	return Number{value: value}
}

func NewNumberFromString(s string) Node {
	return Number{value: datatype.NewNumber(s)}
}

func (Number) NodeType() NodeType { return NodeTypeNumber }

func (Number) IsTerminated() bool { return true }

func (n Number) Value() datatype.Number {
	return n.value
}

func (n Number) MutoString() string {
	if n.value.IsFloat() {
		return fmt.Sprint(n.value.ToFloat())
	}
	return fmt.Sprint(n.value.ToInt())
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
