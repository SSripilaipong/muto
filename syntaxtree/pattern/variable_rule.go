package pattern

type VariableRule struct {
	variableName string
	params       ParamPart
}

func (VariableRule) RulePatternParamType() ParamType {
	return ParamTypeNestedVariableRule
}

func (p VariableRule) VariableName() string {
	return p.variableName
}

func (p VariableRule) ParamPart() ParamPart {
	return p.params
}

func NewVariableRulePattern(variableName string, params ParamPart) VariableRule {
	return VariableRule{variableName: variableName, params: params}
}

func UnsafeRuleParamPatternToVariableRulePattern(p Param) VariableRule {
	return p.(VariableRule)
}

func VariableRulePatternToRulePatternParam(x VariableRule) Param {
	return x
}
