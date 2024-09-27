package syntaxtree

import (
	"github.com/SSripilaipong/muto/core/base/datatype"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Number struct {
	value string
}

func (Number) RuleResultNodeType() stResult.NodeType { return stResult.NodeTypeNumber }

func (Number) ObjectParamType() stResult.ParamType { return stResult.ParamTypeSingle }

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

func UnsafeRuleResultToNumber(p stResult.Node) Number {
	return p.(Number)
}

func UnsafeRuleParamPatternToNumber(p stPattern.Param) Number {
	return p.(Number)
}
