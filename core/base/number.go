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
