package syntaxtree

type RuleResultAnonymousObject struct {
	head      RuleResult
	paramPart ObjectParamPart
}

func NewRuleResultAnonymousObject(head RuleResult, paramPart ObjectParamPart) RuleResultAnonymousObject {
	return RuleResultAnonymousObject{head: head, paramPart: paramPart}
}

func (RuleResultAnonymousObject) RuleResultType() RuleResultType {
	return RuleResultTypeAnonymousObject
}

func (RuleResultAnonymousObject) ObjectParamType() ObjectParamType { return ObjectParamTypeSingle }

func (obj RuleResultAnonymousObject) Head() RuleResult {
	return obj.head
}

func (obj RuleResultAnonymousObject) ParamPart() ObjectParamPart {
	return obj.paramPart
}

type AnonymousObjectHeadType string

const (
	AnonymousObjectHeadTypeNamedObject     AnonymousObjectHeadType = "NAMED_OBJECT"
	AnonymousObjectHeadTypeAnonymousObject AnonymousObjectHeadType = "ANONYMOUS_OBJECT"
)
