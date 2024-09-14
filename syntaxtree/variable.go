package syntaxtree

type Variable struct {
	name string
}

func NewVariable(name string) Variable {
	return Variable{name: name}
}

func (Variable) RuleResultType() RuleResultType {
	return RuleResultTypeVariable
}

func (Variable) ObjectParamType() ObjectParamType { return ObjectParamTypeSingle }

func (Variable) RuleParamPatternType() RuleParamPatternType {
	return RuleParamPatternTypeVariable
}

func (v Variable) Name() string {
	return v.name
}

func UnsafeRuleParamPatternToVariable(p RuleParamPattern) Variable {
	return p.(Variable)
}
