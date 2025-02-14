package base

type PatternParam interface {
	RulePatternParamType() PatternParamType
}

type PatternParamType string

const (
	PatternParamTypeVariable            PatternParamType = "VARIABLE"
	PatternParamTypeBoolean             PatternParamType = "BOOLEAN"
	PatternParamTypeString              PatternParamType = "STRING"
	PatternParamTypeNumber              PatternParamType = "NUMBER"
	PatternParamTypeClass               PatternParamType = "CLASS"
	PatternParamTypeTag                 PatternParamType = "TAG"
	PatternParamTypeNestedNamedRule     PatternParamType = "NESTED_NAMED_RULE"
	PatternParamTypeNestedVariableRule  PatternParamType = "NESTED_VARIABLE_RULE"
	PatternParamTypeNestedAnonymousRule PatternParamType = "NESTED_ANONYMOUS_RULE"
)

func IsPatternParamTypeVariable(p PatternParam) bool {
	return p.RulePatternParamType() == PatternParamTypeVariable
}

func IsPatternParamTypeBoolean(p PatternParam) bool {
	return p.RulePatternParamType() == PatternParamTypeBoolean
}

func IsPatternParamTypeString(p PatternParam) bool {
	return p.RulePatternParamType() == PatternParamTypeString
}

func IsPatternParamTypeNumber(p PatternParam) bool {
	return p.RulePatternParamType() == PatternParamTypeNumber
}

func IsPatternParamTypeTag(p PatternParam) bool {
	return p.RulePatternParamType() == PatternParamTypeTag
}

func IsPatternParamTypeClass(p PatternParam) bool {
	return p.RulePatternParamType() == PatternParamTypeClass
}

func IsPatternParamTypeNestedNamedRule(p PatternParam) bool {
	return p.RulePatternParamType() == PatternParamTypeNestedNamedRule
}

func IsPatternParamTypeNestedVariableRule(p PatternParam) bool {
	return p.RulePatternParamType() == PatternParamTypeNestedVariableRule
}

func IsPatternParamTypeNestedAnonymousRule(p PatternParam) bool {
	return p.RulePatternParamType() == PatternParamTypeNestedAnonymousRule
}

func IsNestedPatternParam(p PatternParam) bool {
	return IsPatternParamTypeNestedNamedRule(p) || IsPatternParamTypeNestedVariableRule(p) || IsPatternParamTypeNestedAnonymousRule(p)
}
