package base

import (
	"fmt"
	"strconv"

	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type String struct {
	value string
}

func NewString(value string) String {
	return String{value: value}
}

func (String) RuleResultNodeType() stResult.NodeType { return stResult.NodeTypeString }

func (String) ObjectParamType() stResult.ParamType { return stResult.ParamTypeSingle }

func (String) RulePatternParamType() stPattern.ParamType {
	return stPattern.ParamTypeString
}

func (s String) Value() string {
	return s.value
}

func (s String) StringValue() string {
	y, err := strconv.Unquote(s.value)
	if err != nil {
		panic(fmt.Errorf("unexpected error: %w", err))
	}
	return y
}

func UnsafeRuleResultToString(r stResult.Node) String {
	return r.(String)
}

func UnsafeRuleParamPatternToString(r stPattern.Param) String {
	return r.(String)
}
