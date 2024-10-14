package base

import "github.com/SSripilaipong/muto/common/optional"

type MutableNode interface {
	Node
	Mutate(mutation Mutation) optional.Of[Node]
}

func IsMutableNode(node Node) bool {
	return IsObjectNode(node) || IsClassNode(node) || IsStructureNode(node)
}

func UnsafeNodeToMutable(x Node) MutableNode {
	return x.(MutableNode)
}
