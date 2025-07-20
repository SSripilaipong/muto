package syntaxtree

import (
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Boolean struct {
	value string
}

func NewBoolean(value string) Boolean {
	return Boolean{value: value}
}

func (Boolean) PatternType() base.PatternType { return base.PatternTypeBoolean }

func (Boolean) RuleResultNodeType() stResult.NodeType { return stResult.NodeTypeBoolean }

func (Boolean) ObjectParamType() stResult.ParamType { return stResult.ParamTypeSingle }

func (b Boolean) Value() string {
	return b.value
}

func (b Boolean) BooleanValue() bool {
	return b.Value() == "true"
}

func UnsafeRuleResultToBoolean(r stResult.Node) Boolean {
	return r.(Boolean)
}

func UnsafePatternToBoolean(p base.Pattern) Boolean {
	return p.(Boolean)
}
