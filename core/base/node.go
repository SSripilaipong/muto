package base

type Node interface {
	NodeType() NodeType
	IsTerminated() bool
}

type NodeType string

const (
	NodeTypeString NodeType = "STRING"
	NodeTypeNumber NodeType = "NUMBER"
	NodeTypeObject NodeType = "OBJECT"
)

func IsObjectNode(node Node) bool {
	return node.NodeType() == NodeTypeObject
}

func IsNumberNode(node Node) bool {
	return node.NodeType() == NodeTypeNumber
}

func IsStringNode(node Node) bool {
	return node.NodeType() == NodeTypeString
}
