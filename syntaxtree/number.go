package syntaxtree

import (
	"github.com/SSripilaipong/muto/core/base/datatype"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

type Number struct {
	value string
}

func (Number) RuleResultType() RuleResultType { return RuleResultTypeNumber }

func (Number) ObjectParamType() ObjectParamType { return ObjectParamTypeSingle }

func (Number) RulePatternParamType() stPattern.ParamType {
	return stPattern.ParamTypeNumber
}

func (n Number) Value() string {
	return n.value
}

func (n Number) NumberValue() datatype.Number {
	return datatype.NewNumber(n.value)
}

func NewNumber(value string) Number {
	return Number{value: value}
}

func IsRuleParamPatternNumber(p stPattern.Param) bool {
	return p.RulePatternParamType() == stPattern.ParamTypeNumber
}

func UnsafeRuleResultToNumber(p RuleResult) Number {
	return p.(Number)
}

func UnsafeRuleParamPatternToNumber(p stPattern.Param) Number {
	return p.(Number)
}
