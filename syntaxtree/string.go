package syntaxtree

import (
	"fmt"
	"strconv"
)

type String struct {
	value string
}

func NewString(value string) String {
	return String{value: value}
}

func (String) RuleResultType() RuleResultType {
	return RuleResultTypeString
}

func (String) RuleParamPatternType() RuleParamPatternType {
	return RuleParamPatternTypeString
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

func UnsafeRuleParamPatternToString(r RuleParamPattern) String {
	return r.(String)
}
