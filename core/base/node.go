package base

import "github.com/SSripilaipong/muto/common/optional"

type Node interface {
	NodeType() NodeType
	MutateAsHead(children []Node, mutation Mutation) optional.Of[Node]
	TopLevelString() string
}

type NodeType string

const (
	NodeTypeString    NodeType = "STRING"
	NodeTypeNumber    NodeType = "NUMBER"
	NodeTypeBoolean   NodeType = "BOOLEAN"
	NodeTypeObject    NodeType = "OBJECT"
	NodeTypeClass     NodeType = "CLASS"
	NodeTypeTag       NodeType = "TAG"
	NodeTypeStructure NodeType = "STRUCTURE"
)

func IsObjectNode(node Node) bool {
	return node.NodeType() == NodeTypeObject
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

func IsTagNode(node Node) bool {
	return node.NodeType() == NodeTypeTag
}

func IsStructureNode(node Node) bool {
	return node.NodeType() == NodeTypeStructure
}

func ToNode[T Node](x T) Node { return x }
