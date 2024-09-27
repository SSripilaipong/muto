package syntaxtree

type Boolean struct {
	value string
}

func NewBoolean(value string) Boolean {
	return Boolean{value: value}
}

func (Boolean) RuleResultType() RuleResultType { return RuleResultTypeBoolean }

func (Boolean) ObjectParamType() ObjectParamType { return ObjectParamTypeSingle }

func (Boolean) RuleParamPatternType() RuleParamPatternType {
	return RuleParamPatternTypeBoolean
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

func UnsafeRuleParamPatternToBoolean(p RuleParamPattern) Boolean {
	return p.(Boolean)
}
