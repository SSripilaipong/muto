package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/data"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func newParamExtractors(params []st.RuleParamPattern) []func(base.Node) optional.Of[*data.Mutation] {
	return slc.Map(newParamExtractor)(params)
}

func newParamExtractor(p st.RuleParamPattern) func(base.Node) optional.Of[*data.Mutation] {
	switch {
	case st.IsRuleParamPatternVariable(p):
		return newVariableParamExtractor(st.UnsafeRuleParamPatternToVariable(p))
	case st.IsRuleParamPatternBoolean(p):
		return newBooleanParamExtractor(st.UnsafeRuleParamPatternToBoolean(p))
	case st.IsRuleParamPatternString(p):
		return newStringParamExtractor(st.UnsafeRuleParamPatternToString(p))
	case st.IsRuleParamPatternNumber(p):
		return newNumberParamExtractor(st.UnsafeRuleParamPatternToNumber(p))
	case st.IsRuleParamPatternNestedNamedRulePattern(p):
		return newNestedNamedRuleExtractor(st.UnsafeRuleParamPatternToNamedRulePattern(p))
	case st.IsRuleParamPatternNestedVariableRulePattern(p):
		return newNestedVariableRuleExtractor(st.UnsafeRuleParamPatternToVariableRulePattern(p))
	case st.IsRuleParamPatternNestedAnonymousRulePattern(p):
		return newNestedAnonymousRuleExtractor(st.UnsafeRuleParamPatternToAnonymousRulePattern(p))
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
