package syntaxtree

type RuleParamPattern interface {
	RuleParamPatternType() RuleParamPatternType
}

type RuleParamPatternType string

const (
	RuleParamPatternTypeVariable RuleParamPatternType = "VARIABLE"
	RuleParamPatternTypeString   RuleParamPatternType = "STRING"
)

func IsRuleParamPatternVariable(p RuleParamPattern) bool {
	return p.RuleParamPatternType() == RuleParamPatternTypeVariable
}

func IsRuleParamPatternString(p RuleParamPattern) bool {
	return p.RuleParamPatternType() == RuleParamPatternTypeString
}
