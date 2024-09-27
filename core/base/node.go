package base

import "github.com/SSripilaipong/muto/common/optional"

type Node interface {
	NodeType() NodeType
	MutateAsHead(children []Node, mutation Mutation) optional.Of[Node]
}

type NodeType string

const (
	NodeTypeString  NodeType = "STRING"
	NodeTypeNumber  NodeType = "NUMBER"
	NodeTypeBoolean NodeType = "BOOLEAN"
	NodeTypeObject  NodeType = "OBJECT"
	NodeTypeClass   NodeType = "CLASS"
)

func IsObjectNode(node Node) bool {
	return node.NodeType() == NodeTypeObject
}

func IsMutableNode(node Node) bool {
	return IsObjectNode(node) || IsClassNode(node)
}

func IsClassNode(node Node) bool {
	return node.NodeType() == NodeTypeClass
}

func IsBooleanNode(node Node) bool {
	return node.NodeType() == NodeTypeBoolean
}

func IsNumberNode(node Node) bool {
	return node.NodeType() == NodeTypeNumber
}

func IsStringNode(node Node) bool {
	return node.NodeType() == NodeTypeString
}
