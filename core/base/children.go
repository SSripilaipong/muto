package base

import (
	"slices"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
)

func mutateChildren(params ParamChain, mutation NameWiseMutation) optional.Of[ParamChain] {
	children := slices.Clone(params.DirectParams())

	for i, child := range children {
		if IsMutableNode(child) {
			childObj := UnsafeNodeToMutable(child)
			if newChild := childObj.Mutate(mutation); newChild.IsNotEmpty() {
				children[i] = newChild.Value()
				return optional.Value(NewParamChain(slc.Pure(children)))
			}
		}
	}

	return optional.Empty[ParamChain]()
}
