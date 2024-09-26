package base

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/optional"
)

type String struct {
	value string
}

func NewString(value string) String {
	return String{value: value}
}

func (String) NodeType() NodeType { return NodeTypeString }

func (s String) MutateAsHead(children []Node, mutation Mutation) optional.Of[Node] {
	if len(children) > 0 {
		newChildren := mutateChildren(children, mutation)
		if newChildren.IsEmpty() {
			return optional.Empty[Node]()
		}
		return optional.Value[Node](NewObject(s, newChildren.Value()))
	}
	return optional.Value[Node](s)
}

func (s String) Value() string {
	return s.value
}

func (s String) String() string {
	return fmt.Sprintf("%#v", s.value)
}

func UnsafeNodeToString(n Node) String {
	return n.(String)
}

func StringToNode(s String) Node {
	return s
}
