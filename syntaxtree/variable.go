package syntaxtree

import stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"

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

func (Variable) RulePatternParamType() stPattern.ParamType {
	return stPattern.ParamTypeVariable
}

func (v Variable) Name() string {
	return v.name
}

func UnsafeRuleResultToVariable(p RuleResult) Variable {
	return p.(Variable)
}

func UnsafeRuleParamPatternToVariable(p stPattern.Param) Variable {
	return p.(Variable)
}
