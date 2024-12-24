package base

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/optional"
)

type Tag struct {
	name string
}

// NewTag expects name without leading dot
func NewTag(name string) Node {
	return Tag{name: name}
}

func (Tag) NodeType() NodeType { return NodeTypeTag }

func (t Tag) MutateAsHead(params ParamChain, mutation NameWiseMutation) optional.Of[Node] {
	children := params.DirectParams()
	if len(children) > 0 {
		newChildren := mutateChildren(params, mutation)
		if newChildren.IsEmpty() {
			return optional.Empty[Node]()
		}
		return optional.Value[Node](NewObject(t, newChildren.Value()))
	}
	return optional.Value[Node](t)
}

func (t Tag) Name() string {
	return t.name
}

func (t Tag) MutoString() string {
	return fmt.Sprintf(".%s", t.Name())
}

func (t Tag) TopLevelString() string {
	return t.String()
}

func (t Tag) String() string {
	return t.MutoString()
}

func UnsafeNodeToTag(n Node) Tag {
	return n.(Tag)
}
