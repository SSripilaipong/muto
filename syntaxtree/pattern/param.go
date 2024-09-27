package pattern

type Param interface {
	RulePatternParamType() ParamType
}

type ParamType string

const (
	ParamTypeVariable            ParamType = "VARIABLE"
	ParamTypeBoolean             ParamType = "BOOLEAN"
	ParamTypeString              ParamType = "STRING"
	ParamTypeNumber              ParamType = "NUMBER"
	ParamTypeNestedNamedRule     ParamType = "NESTED_NAMED_RULE"
	ParamTypeNestedVariableRule  ParamType = "NESTED_VARIABLE_RULE"
	ParamTypeNestedAnonymousRule ParamType = "NESTED_ANONYMOUS_RULE"
)

func IsParamTypeVariable(p Param) bool {
	return p.RulePatternParamType() == ParamTypeVariable
}

func IsParamTypeBoolean(p Param) bool {
	return p.RulePatternParamType() == ParamTypeBoolean
}

func IsParamTypeString(p Param) bool {
	return p.RulePatternParamType() == ParamTypeString
}

func IsParamTypeNestedNamedRule(p Param) bool {
	return p.RulePatternParamType() == ParamTypeNestedNamedRule
}

func IsParamTypeNestedVariableRule(p Param) bool {
	return p.RulePatternParamType() == ParamTypeNestedVariableRule
}

func IsParamTypeNestedAnonymousRule(p Param) bool {
	return p.RulePatternParamType() == ParamTypeNestedAnonymousRule
}
