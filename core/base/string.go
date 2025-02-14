package base

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
)

type String struct {
	value string
}

func NewString(value string) String {
	return String{value: value}
}

var NewStringObject = fn.Compose(NewPrimitiveObject, NewString)

func (String) NodeType() NodeType { return NodeTypeString }

func (s String) MutateAsHead(params ParamChain, mutation NameWiseMutation) optional.Of[Node] {
	newParams := mutateParamChain(params, mutation)
	if newParams.IsNotEmpty() {
		return optional.Value[Node](NewCompoundObject(s, newParams.Value()))
	}
	return optional.Empty[Node]()
}

func (s String) Value() string {
	return s.value
}

func (s String) TopLevelString() string {
	return s.String()
}

func (s String) String() string {
	return fmt.Sprintf("%#v", s.value)
}

func (s String) MutoString() string {
	return s.Value()
}

func UnsafeNodeToString(n Node) String {
	return n.(String)
}

func StringToNode(s String) Node {
	return s
}
