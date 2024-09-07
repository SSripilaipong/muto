package syntaxtree

type RuleResultAnonymousObject struct {
	head   RuleResult
	params []ObjectParam
}

func NewRuleResultAnonymousObject(head RuleResult, params []ObjectParam) RuleResultAnonymousObject {
	return RuleResultAnonymousObject{head: head, params: params}
}

func (RuleResultAnonymousObject) RuleResultType() RuleResultType {
	return RuleResultTypeAnonymousObject
}

func (obj RuleResultAnonymousObject) Head() RuleResult {
	return obj.head
}

func (obj RuleResultAnonymousObject) Params() []ObjectParam {
	return obj.params
}

type AnonymousObjectHeadType string

const (
	AnonymousObjectHeadTypeNamedObject     AnonymousObjectHeadType = "NAMED_OBJECT"
	AnonymousObjectHeadTypeAnonymousObject AnonymousObjectHeadType = "ANONYMOUS_OBJECT"
)
