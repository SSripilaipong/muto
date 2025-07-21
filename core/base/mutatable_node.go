package base

import "github.com/SSripilaipong/go-common/optional"

type MutableNode interface {
	Node
	Mutate() optional.Of[Node]
}

func IsMutableNode(node Node) bool {
	return IsObjectNode(node) || IsStructureNode(node)
}

func UnsafeNodeToMutable(x Node) MutableNode {
	return x.(MutableNode)
}
