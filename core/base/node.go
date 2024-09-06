package base

type Node interface {
	NodeType() NodeType
	IsTerminationConfirmed() bool
}

type NodeType string

const (
	NodeTypeString      NodeType = "STRING"
	NodeTypeNumber      NodeType = "NUMBER"
	NodeTypeNamedObject NodeType = "NAMED_OBJECT"
)

func IsNamedObjectNode(node Node) bool {
	return node.NodeType() == NodeTypeNamedObject
}

func IsNumberNode(node Node) bool {
	return node.NodeType() == NodeTypeNumber
}

func IsStringNode(node Node) bool {
	return node.NodeType() == NodeTypeString
}
