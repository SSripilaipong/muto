package global

import "github.com/SSripilaipong/muto/core/base"

func mutateUntilTerminated(node base.Node) base.Node {
	if !base.IsMutableNode(node) {
		return node
	}

	if result, isMutated := base.UnsafeNodeToMutable(node).Mutate().Return(); isMutated {
		return mutateUntilTerminated(result)
	}
	return node
}
