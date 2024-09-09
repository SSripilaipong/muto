package syntaxtree

type RuleParamPattern interface {
	RuleParamPatternType() RuleParamPatternType
}

type RuleParamPatternType string

const (
	RuleParamPatternTypeVariable                  RuleParamPatternType = "VARIABLE"
	RuleParamPatternTypeString                    RuleParamPatternType = "STRING"
	RuleParamPatternTypeNumber                    RuleParamPatternType = "NUMBER"
	RuleParamPatternTypeNestedNamedRulePattern    RuleParamPatternType = "NESTED_NAMED_RULE_PATTERN"
	RuleParamPatternTypeNestedVariableRulePattern RuleParamPatternType = "NESTED_VARIABLE_RULE_PATTERN"
)

func IsRuleParamPatternVariable(p RuleParamPattern) bool {
	return p.RuleParamPatternType() == RuleParamPatternTypeVariable
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
