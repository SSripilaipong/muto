package base

// MutateUntilTerminated mutates the node until it becomes immutable or a mutation returns no result.
func MutateUntilTerminated(node Node) Node {
	for IsMutableNode(node) {
		next, ok := UnsafeNodeToMutable(node).Mutate().Return()
		if !ok {
			break
		}
		node = next
	}
	return node
}
