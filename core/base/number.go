package base

import (
	"fmt"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base/datatype"
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

func (n Number) MutateAsHead(params ParamChain) optional.Of[Node] {
	newChildren := MutateParamChain(params)
	if newChildren.IsNotEmpty() {
		return optional.Value[Node](NewCompoundObject(n, newChildren.Value()))
	}
	return optional.Empty[Node]()
}

func (n Number) Value() datatype.Number {
	return n.value
}

func (n Number) MutoString() string {
	return n.Value().SimpleString()
}

func (n Number) TopLevelString() string {
	return n.String()
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
