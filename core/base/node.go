package base

type Node interface {
	NodeType() NodeType
	IsTerminationConfirmed() bool
}

type NodeType string

const (
	NodeTypeString          NodeType = "STRING"
	NodeTypeNumber          NodeType = "NUMBER"
	NodeTypeNamedObject     NodeType = "NAMED_OBJECT"
	NodeTypeAnonymousObject NodeType = "ANONYMOUS_OBJECT"
	NodeTypeNamedClass      NodeType = "NAMED_CLASS"
)

func IsObjectNode(node Node) bool {
	return IsNamedObjectNode(node) || IsAnonymousObjectNode(node)
}

func IsMutableNode(node Node) bool {
	return IsNamedObjectNode(node) || IsAnonymousObjectNode(node) || IsNamedClassNode(node)
}

func IsNamedClassNode(node Node) bool {
	return node.NodeType() == NodeTypeNamedClass
}

func IsNamedObjectNode(node Node) bool {
	return node.NodeType() == NodeTypeNamedObject
}

func IsAnonymousObjectNode(node Node) bool {
	return node.NodeType() == NodeTypeAnonymousObject
}

func IsNumberNode(node Node) bool {
	return node.NodeType() == NodeTypeNumber
}

func IsStringNode(node Node) bool {
	return node.NodeType() == NodeTypeString
}
