package result

type Node interface {
	RuleResultNodeType() NodeType
}

type NodeType string

const (
	NodeTypeBoolean          NodeType = "BOOLEAN"
	NodeTypeString           NodeType = "STRING"
	NodeTypeNumber           NodeType = "NUMBER"
	NodeTypeClass            NodeType = "CLASS"
	NodeTypeNamedObject      NodeType = "NAMED_OBJECT"
	NodeTypeObject           NodeType = "OBJECT"
	NodeTypeVariable         NodeType = "VARIABLE"
	NodeTypeVariadicVariable NodeType = "VARIADIC_VARIABLE"
)

func IsNodeTypeBoolean(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeBoolean
}

func IsNodeTypeString(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeString
}

func IsNodeTypeNumber(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeNumber
}

func IsNodeTypeClass(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeClass
}

func IsNodeTypeAnonymousObject(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeObject
}

func IsNodeTypeNamedObject(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeNamedObject
}

func IsNodeTypeVariable(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeVariable
}
