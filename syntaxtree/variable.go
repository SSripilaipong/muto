package syntaxtree

import (
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Variable struct {
	name string
}

func NewVariable(name string) Variable {
	return Variable{name: name}
}

func (Variable) PatternType() base.PatternType { return base.PatternTypeVariable }

func (Variable) RuleResultNodeType() stResult.NodeType {
	return stResult.NodeTypeVariable
}

func (Variable) ObjectParamType() stResult.ParamType { return stResult.ParamTypeSingle }

func (Variable) NonObjectNode() {}

func (v Variable) Name() string {
	return v.name
}

func UnsafeRuleResultToVariable(p stResult.Node) Variable {
	return p.(Variable)
}

func UnsafePatternToVariable(p base.Pattern) Variable {
	return p.(Variable)
}
