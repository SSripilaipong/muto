package syntaxtree

type RuleResultNamedObject struct {
	objectName string
	params     []ObjectParam
}

func NewRuleResultNamedObject(objectName string, params []ObjectParam) RuleResultNamedObject {
	return RuleResultNamedObject{objectName: objectName, params: params}
}

func (RuleResultNamedObject) RuleResultType() RuleResultType { return RuleResultTypeNamedObject }

func (RuleResultNamedObject) AnonymousObjectHeadType() AnonymousObjectHeadType {
	return AnonymousObjectHeadTypeNamedObject
}

func (obj RuleResultNamedObject) ObjectName() string {
	return obj.objectName
}

func (obj RuleResultNamedObject) Params() []ObjectParam {
	return obj.params
}
