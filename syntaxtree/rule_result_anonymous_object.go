package syntaxtree

type RuleResultAnonymousObject struct {
	head   AnonymousObjectHead
	params []ObjectParam
}

func NewRuleResultAnonymousObject(head AnonymousObjectHead, params []ObjectParam) RuleResultAnonymousObject {
	return RuleResultAnonymousObject{head: head, params: params}
}

func (RuleResultAnonymousObject) RuleResultType() RuleResultType {
	return RuleResultTypeAnonymousObject
}

func (RuleResultAnonymousObject) AnonymousObjectHeadType() AnonymousObjectHeadType {
	return AnonymousObjectHeadTypeAnonymousObject
}

func (obj RuleResultAnonymousObject) Params() []ObjectParam {
	return obj.params
}

type AnonymousObjectHeadType string

const (
	AnonymousObjectHeadTypeNamedObject     AnonymousObjectHeadType = "NAMED_OBJECT"
	AnonymousObjectHeadTypeAnonymousObject AnonymousObjectHeadType = "ANONYMOUS_OBJECT"
)

type AnonymousObjectHead interface {
	AnonymousObjectHeadType() AnonymousObjectHeadType
}
