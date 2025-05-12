package syntaxtree

import (
	"github.com/SSripilaipong/muto/core/base/datatype"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Number struct {
	value string
}

func (Number) PatternType() base.PatternType { return base.PatternTypeNumber }

func (Number) RuleResultNodeType() stResult.NodeType { return stResult.NodeTypeNumber }

func (Number) ObjectParamType() stResult.ParamType { return stResult.ParamTypeSingle }

func (Number) NonObjectNode() {}

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

func UnsafePatternToNumber(p base.Pattern) Number {
	return p.(Number)
}
