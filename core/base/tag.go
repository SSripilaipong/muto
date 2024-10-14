package base

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/optional"
)

type Tag struct {
	name string
}

func NewTag(name string) Node {
	return Tag{name: name}
}

func (Tag) NodeType() NodeType { return NodeTypeTag }

func (t Tag) MutateAsHead(children []Node, mutation Mutation) optional.Of[Node] {
	if len(children) > 0 {
		newChildren := mutateChildren(children, mutation)
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
