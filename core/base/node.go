package base

type Node interface {
	NodeType() NodeType
}

type NodeType string

const (
	NodeTypeString   NodeType = "STRING"
	NodeTypeNumber   NodeType = "NUMBER"
	NodeTypeObject   NodeType = "OBJECT"
	NodeTypeVariable NodeType = "VARIABLE"
)

func IsObjectNode(node Node) bool {
	return node.NodeType() == NodeTypeObject
}

func IsVariableNode(node Node) bool {
	return node.NodeType() == NodeTypeVariable
}
