package syntaxtree

import stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"

type Boolean struct {
	value string
}

func NewBoolean(value string) Boolean {
	return Boolean{value: value}
}

func (Boolean) RuleResultType() RuleResultType { return RuleResultTypeBoolean }

func (Boolean) ObjectParamType() ObjectParamType { return ObjectParamTypeSingle }

func (Boolean) RulePatternParamType() stPattern.ParamType {
	return stPattern.ParamTypeBoolean
}

func (s Boolean) Value() string {
	return s.value
}

func (s Boolean) BooleanValue() bool {
	return s.Value() == "true"
}

func UnsafeRuleResultToBoolean(r RuleResult) Boolean {
	return r.(Boolean)
}

func UnsafeRuleParamPatternToBoolean(p stPattern.Param) Boolean {
	return p.(Boolean)
}
