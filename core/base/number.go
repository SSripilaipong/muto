package base

import "phi-lang/core/base/datatype"

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

func UnsafeNodeToNumber(n Node) Number {
	return n.(Number)
}
