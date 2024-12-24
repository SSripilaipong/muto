package base

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/optional"
)

type Boolean struct {
	value bool
}

func NewBoolean(x bool) Boolean {
	return Boolean{value: x}
}

func (Boolean) NodeType() NodeType { return NodeTypeBoolean }

func (b Boolean) MutateAsHead(params ParamChain, mutation NameWiseMutation) optional.Of[Node] {
	children := params.DirectParams()
	if len(children) > 0 {
		newChildren := mutateChildren(children, mutation)
		if newChildren.IsEmpty() {
			return optional.Empty[Node]()
		}
		return optional.Value[Node](NewObject(b, newChildren.Value()))
	}
	return optional.Value[Node](b)
}

func (b Boolean) Value() bool {
	return b.value
}

func (b Boolean) TopLevelString() string {
	return b.String()
}

func (b Boolean) String() string {
	return fmt.Sprintf("%v", b.value)
}

func (b Boolean) MutoString() string {
	return b.String()
}

func UnsafeNodeToBoolean(n Node) Boolean {
	return n.(Boolean)
}
