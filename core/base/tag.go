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
	newChildren := mutateParamChain(params, mutation)
	if newChildren.IsNotEmpty() {
		return optional.Value[Node](NewCompoundObject(t, newChildren.Value()))
	}
	return optional.Empty[Node]()
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
