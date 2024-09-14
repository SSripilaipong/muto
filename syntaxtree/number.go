package syntaxtree

import "muto/core/base/datatype"

type Number struct {
	value string
}

func (Number) RuleResultType() RuleResultType { return RuleResultTypeNumber }

func (Number) ObjectParamType() ObjectParamType { return ObjectParamTypeSingle }

func (Number) RuleParamPatternType() RuleParamPatternType { return RuleParamPatternTypeNumber }

func (n Number) Value() string {
	return n.value
}

func (n Number) NumberValue() datatype.Number {
	return datatype.NewNumber(n.value)
}

func NewNumber(value string) Number {
	return Number{value: value}
}

func IsRuleParamPatternNumber(p RuleParamPattern) bool {
	return p.RuleParamPatternType() == RuleParamPatternTypeNumber
}

func UnsafeRuleParamPatternToNumber(p RuleParamPattern) Number {
	return p.(Number)
}
