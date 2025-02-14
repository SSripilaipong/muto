package base

import (
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Variable struct {
	name string
}

func NewVariable(name string) Variable {
	return Variable{name: name}
}

func (Variable) RuleResultNodeType() stResult.NodeType {
	return stResult.NodeTypeVariable
}

func (Variable) ObjectParamType() stResult.ParamType { return stResult.ParamTypeSingle }

func (Variable) RulePatternParamType() PatternParamType {
	return PatternParamTypeVariable
}

func (v Variable) Name() string {
	return v.name
}

func UnsafeRuleResultToVariable(p stResult.Node) Variable {
	return p.(Variable)
}

func UnsafeRuleParamPatternToVariable(p PatternParam) Variable {
	return p.(Variable)
}
