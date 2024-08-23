package syntaxtree

type RulePattern struct {
	objectName string
	params     []RuleParamPattern
}

func NewRulePattern(objectName string, params []RuleParamPattern) RulePattern {
	return RulePattern{objectName: objectName, params: params}
}

type RuleParamPattern interface{}
