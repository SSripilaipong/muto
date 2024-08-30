package syntaxtree

type RulePattern struct {
	objectName string
	params     []RuleParamPattern
}

func (RulePattern) RuleParamPatternType() RuleParamPatternType {
	return RuleParamPatternTypeNestedRulePattern
}

func (p RulePattern) ObjectName() string {
	return p.objectName
}

func (p RulePattern) NParams() uint8 {
	return uint8(len(p.Params()))
}

func (p RulePattern) Params() []RuleParamPattern {
	return p.params
}

func NewRulePattern(objectName string, params []RuleParamPattern) RulePattern {
	return RulePattern{objectName: objectName, params: params}
}

func ParamsOfRulePattern(p RulePattern) []RuleParamPattern {
	return p.Params()
}

func UnsafeRuleParamPatternToRulePattern(p RuleParamPattern) RulePattern {
	return p.(RulePattern)
}
