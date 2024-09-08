package syntaxtree

type RulePattern struct {
	objectName string
	params     RulePatternParamPart
}

func (RulePattern) RuleParamPatternType() RuleParamPatternType {
	return RuleParamPatternTypeNestedRulePattern
}

func (p RulePattern) ObjectName() string {
	return p.objectName
}

func (p RulePattern) CheckNParams(n int) bool {
	return p.ParamPart().CheckNParams(n)
}

func (p RulePattern) ParamPart() RulePatternParamPart {
	return p.params
}

func NewRulePattern(objectName string, params RulePatternParamPart) RulePattern {
	return RulePattern{objectName: objectName, params: params}
}

func ParamPartOfRulePattern(p RulePattern) RulePatternParamPart {
	return p.ParamPart()
}

func UnsafeRuleParamPatternToRulePattern(p RuleParamPattern) RulePattern {
	return p.(RulePattern)
}

type RulePatternParamPart interface {
	RulePatternParamPartType() RulePatternParamPartType
	CheckNParams(n int) bool
}

type RulePatternParamPartType string

const (
	RulePatternParamPartTypeFixed    RulePatternParamPartType = "FIXED"
	RulePatternParamPartTypeVariadic RulePatternParamPartType = "VARIADIC"
)

func IsRulePatternParamPartTypeFixed(p RulePatternParamPart) bool {
	return p.RulePatternParamPartType() == RulePatternParamPartTypeFixed
}

func IsRulePatternParamPartTypeVariadic(p RulePatternParamPart) bool {
	return p.RulePatternParamPartType() == RulePatternParamPartTypeVariadic
}
