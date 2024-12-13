package base

import (
	"slices"

	"github.com/SSripilaipong/muto/common/optional"
)

func mutateChildren(children []Node, mutation NameWiseMutation) optional.Of[[]Node] {
	children = slices.Clone(children)

	for i, child := range children {
		if IsMutableNode(child) {
			childObj := UnsafeNodeToMutable(child)
			if newChild := childObj.Mutate(mutation); newChild.IsNotEmpty() {
				children[i] = newChild.Value()
				return optional.Value(children)
			}
		}
	}

	return optional.Empty[[]Node]()
}
