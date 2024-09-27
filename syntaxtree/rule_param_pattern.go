package syntaxtree

type RuleParamPattern interface {
	RuleParamPatternType() RuleParamPatternType
}

type RuleParamPatternType string

const (
	RuleParamPatternTypeVariable                   RuleParamPatternType = "VARIABLE"
	RuleParamPatternTypeBoolean                    RuleParamPatternType = "BOOLEAN"
	RuleParamPatternTypeString                     RuleParamPatternType = "STRING"
	RuleParamPatternTypeNumber                     RuleParamPatternType = "NUMBER"
	RuleParamPatternTypeNestedNamedRulePattern     RuleParamPatternType = "NESTED_NAMED_RULE_PATTERN"
	RuleParamPatternTypeNestedVariableRulePattern  RuleParamPatternType = "NESTED_VARIABLE_RULE_PATTERN"
	RuleParamPatternTypeNestedAnonymousRulePattern RuleParamPatternType = "NESTED_ANONYMOUS_RULE_PATTERN"
)

func IsRuleParamPatternVariable(p RuleParamPattern) bool {
	return p.RuleParamPatternType() == RuleParamPatternTypeVariable
}

func IsRuleParamPatternBoolean(p RuleParamPattern) bool {
	return p.RuleParamPatternType() == RuleParamPatternTypeBoolean
}

func IsRuleParamPatternString(p RuleParamPattern) bool {
	return p.RuleParamPatternType() == RuleParamPatternTypeString
}

func IsRuleParamPatternNestedNamedRulePattern(p RuleParamPattern) bool {
	return p.RuleParamPatternType() == RuleParamPatternTypeNestedNamedRulePattern
}

func IsRuleParamPatternNestedVariableRulePattern(p RuleParamPattern) bool {
	return p.RuleParamPatternType() == RuleParamPatternTypeNestedVariableRulePattern
}

func IsRuleParamPatternNestedAnonymousRulePattern(p RuleParamPattern) bool {
	return p.RuleParamPatternType() == RuleParamPatternTypeNestedAnonymousRulePattern
}
