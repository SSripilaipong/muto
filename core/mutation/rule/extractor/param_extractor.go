package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/data"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func newParamExtractors(params []stPattern.Param) []func(base.Node) optional.Of[*data.Mutation] {
	return slc.Map(newParamExtractor)(params)
}

func newParamExtractor(p stPattern.Param) func(base.Node) optional.Of[*data.Mutation] {
	switch {
	case stPattern.IsParamTypeVariable(p):
		return newVariableParamExtractor(st.UnsafeRuleParamPatternToVariable(p))
	case stPattern.IsParamTypeBoolean(p):
		return newBooleanParamExtractor(st.UnsafeRuleParamPatternToBoolean(p))
	case stPattern.IsParamTypeString(p):
		return newStringParamExtractor(st.UnsafeRuleParamPatternToString(p))
	case stPattern.IsParamTypeNumber(p):
		return newNumberParamExtractor(st.UnsafeRuleParamPatternToNumber(p))
	case stPattern.IsParamTypeTag(p):
		return newTagParamExtractor(st.UnsafeRuleParamPatternToTag(p))
	case stPattern.IsParamTypeNestedNamedRule(p):
		return newNestedNamedRuleExtractor(stPattern.UnsafeParamToNamedRule(p))
	case stPattern.IsParamTypeNestedVariableRule(p):
		return newNestedVariableRuleExtractor(stPattern.UnsafeRuleParamPatternToVariableRulePattern(p))
	case stPattern.IsParamTypeNestedAnonymousRule(p):
		return newNestedAnonymousRuleExtractor(stPattern.UnsafeParamToAnonymousRule(p))
	}
	panic("not implemented")
}

func newVariableParamExtractor(v st.Variable) func(base.Node) optional.Of[*data.Mutation] {
	return func(x base.Node) optional.Of[*data.Mutation] {
		return optional.Value(data.NewMutationWithVariableMapping(data.NewVariableMapping(v.Name(), x)))
	}
}

func newBooleanParamExtractor(v st.Boolean) func(base.Node) optional.Of[*data.Mutation] {
	value := v.BooleanValue()
	return func(x base.Node) optional.Of[*data.Mutation] {
		if base.IsBooleanNode(x) && base.UnsafeNodeToBoolean(x).Value() == value {
			return optional.Value(data.NewMutation())
		}
		return optional.Empty[*data.Mutation]()
	}
}

func newStringParamExtractor(v st.String) func(base.Node) optional.Of[*data.Mutation] {
	value := v.StringValue()
	return func(x base.Node) optional.Of[*data.Mutation] {
		if base.IsStringNode(x) && base.UnsafeNodeToString(x).Value() == value {
			return optional.Value(data.NewMutation())
		}
		return optional.Empty[*data.Mutation]()
	}
}

func newNumberParamExtractor(v st.Number) func(base.Node) optional.Of[*data.Mutation] {
	value := v.NumberValue()
	return func(x base.Node) optional.Of[*data.Mutation] {
		if base.IsNumberNode(x) && base.UnsafeNodeToNumber(x).Value() == value {
			return optional.Value(data.NewMutation())
		}
		return optional.Empty[*data.Mutation]()
	}
}

func newTagParamExtractor(v st.Tag) func(base.Node) optional.Of[*data.Mutation] {
	name := v.Name()
	return func(x base.Node) optional.Of[*data.Mutation] {
		if base.IsTagNode(x) && base.UnsafeNodeToTag(x).Name() == name {
			return optional.Value(data.NewMutation())
		}
		return optional.Empty[*data.Mutation]()
	}
}
