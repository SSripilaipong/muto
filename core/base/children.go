package base

import (
	"slices"

	"github.com/SSripilaipong/muto/common/optional"
)

func MutateParamChain(params ParamChain) optional.Of[ParamChain] {
	nodesList := params.All()

	for i := range nodesList {
		newNodes, ok := mutateChildren(nodesList[i]).Return()
		if ok {
			nodesList[i] = newNodes
			return optional.Value(NewParamChain(nodesList))
		}
	}
	return optional.Empty[ParamChain]()
}

func mutateChildren(children []Node) optional.Of[[]Node] {
	children = slices.Clone(children)
	for i, child := range children {
		if IsMutableNode(child) {
			childObj := UnsafeNodeToMutable(child)
			if newChild := childObj.Mutate(); newChild.IsNotEmpty() {
				children[i] = newChild.Value()
				return optional.Value(children)
			}
		}
	}

	return optional.Empty[[]Node]()
}
