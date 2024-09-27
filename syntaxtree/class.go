package syntaxtree

import (
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Class struct {
	name string
}

func (Class) RuleResultNodeType() stResult.NodeType { return stResult.NodeTypeClass }

func (Class) ObjectParamType() stResult.ParamType { return stResult.ParamTypeSingle }

func (Class) RulePatternParamType() stPattern.ParamType {
	return stPattern.ParamTypeClass
}

func (c Class) Value() string {
	return c.Name()
}

func (c Class) Name() string {
	return c.name
}

func NewClass(name string) Class {
	return Class{name: name}
}

func UnsafeRuleResultToClass(p stResult.Node) Class {
	return p.(Class)
}

func UnsafeRuleParamPatternToClass(p stPattern.Param) Class {
	return p.(Class)
}
