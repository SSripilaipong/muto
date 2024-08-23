package syntaxtree

type RuleResult interface{}

type RuleResultObject struct {
	objectName string
	params     []ObjectParam
}

func NewRuleResultObject(objectName string, params []ObjectParam) RuleResultObject {
	return RuleResultObject{objectName: objectName, params: params}
}

type ObjectParam interface{}
