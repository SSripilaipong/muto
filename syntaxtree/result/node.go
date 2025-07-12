package result

type Node interface {
	RuleResultNodeType() NodeType
	ObjectParamType() ParamType
}

type NodeType string

const (
	NodeTypeBoolean          NodeType = "BOOLEAN"
	NodeTypeString           NodeType = "STRING"
	NodeTypeRune             NodeType = "RUNE"
	NodeTypeNumber           NodeType = "NUMBER"
	NodeTypeClass            NodeType = "CLASS"
	NodeTypeTag              NodeType = "TAG"
	NodeTypeObject           NodeType = "OBJECT"
	NodeTypeNakedObject      NodeType = "NAKED_OBJECT"
	NodeTypeVariable         NodeType = "VARIABLE"
	NodeTypeVariadicVariable NodeType = "VARIADIC_VARIABLE"
	NodeTypeStructure        NodeType = "STRUCTURE"
	NodeTypeReconstructor    NodeType = "RECONSTRUCTOR"
)

func IsNodeTypeBoolean(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeBoolean
}

func IsNodeTypeString(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeString
}

func IsNodeTypeRune(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeRune
}

func IsNodeTypeNumber(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeNumber
}

func IsNodeTypeClass(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeClass
}

func IsNodeTypeTag(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeTag
}

func IsNodeTypeObject(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeObject
}

func IsNodeTypeNakedObject(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeNakedObject
}

func IsNodeTypeVariable(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeVariable
}

func IsNodeTypeStructure(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeStructure
}

func IsNodeTypeReconstructor(r Node) bool {
	return r.RuleResultNodeType() == NodeTypeReconstructor
}

func ToNode[T Node](x T) Node { return x }
