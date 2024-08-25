package syntaxtree

type RuleParamPattern interface {
	RuleParamPatternType() RuleParamPatternType
}

type RuleParamPatternType string

const (
	RuleParamPatternTypeVariable RuleParamPatternType = "VARIABLE"
)

func IsRuleParamPatternVariable(p RuleParamPattern) bool {
	return p.RuleParamPatternType() == RuleParamPatternTypeVariable
}
