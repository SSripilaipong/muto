package base

import (
	"slices"

	"github.com/SSripilaipong/muto/common/optional"
)

func mutateParamChain(params ParamChain, mutation NameWiseMutation) optional.Of[ParamChain] {
	nodesList := params.All()

	for i := len(nodesList) - 1; i >= 0; i-- {
		newNodes, ok := mutateChildren(nodesList[i], mutation).Return()
		if ok {
			nodesList[i] = newNodes
			return optional.Value(NewParamChain(nodesList))
		}
	}
	return optional.Empty[ParamChain]()
}

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
