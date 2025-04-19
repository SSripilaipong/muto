package base

import (
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Boolean struct {
	value string
}

func NewBoolean(value string) Boolean {
	return Boolean{value: value}
}

func (Boolean) PatternType() PatternType { return PatternTypeBoolean }

func (Boolean) RuleResultNodeType() stResult.NodeType { return stResult.NodeTypeBoolean }

func (Boolean) ObjectParamType() stResult.ParamType { return stResult.ParamTypeSingle }

func (Boolean) NonObjectNode() {}

func (b Boolean) Value() string {
	return b.value
}

func (b Boolean) BooleanValue() bool {
	return b.Value() == "true"
}

func UnsafeRuleResultToBoolean(r stResult.Node) Boolean {
	return r.(Boolean)
}

func UnsafePatternToBoolean(p Pattern) Boolean {
	return p.(Boolean)
}
