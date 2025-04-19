package result

type NonObjectNode interface {
	Node
	NonObjectNode()
}
