package extractor

import (
	"muto/common/optional"
	"muto/common/slc"
	"muto/core/base"
	"muto/core/mutation/rule/data"
	st "muto/syntaxtree"
)

func newParamExtractors(params []st.RuleParamPattern) []func(base.Node) optional.Of[*data.Mutation] {
	return slc.Map(newParamExtractor)(params)
}

func newParamExtractor(p st.RuleParamPattern) func(base.Node) optional.Of[*data.Mutation] {
	if st.IsRuleParamPatternVariable(p) {
		return newVariableParamExtractor(st.UnsafeRuleParamPatternToVariable(p))
	} else if st.IsRuleParamPatternString(p) {
		return newStringParamExtractor(st.UnsafeRuleParamPatternToString(p))
	} else if st.IsRuleParamPatternNumber(p) {
		return newNumberParamExtractor(st.UnsafeRuleParamPatternToNumber(p))
	} else if st.IsRuleParamPatternNestedNamedRulePattern(p) {
		return newNestedNamedRuleExtractor(st.UnsafeRuleParamPatternToNamedRulePattern(p))
	} else if st.IsRuleParamPatternNestedVariableRulePattern(p) {
		return newNestedVariableRuleExtractor(st.UnsafeRuleParamPatternToVariableRulePattern(p))
	} else if st.IsRuleParamPatternNestedAnonymousRulePattern(p) {
		return newNestedAnonymousRuleExtractor(st.UnsafeRuleParamPatternToAnonymousRulePattern(p))
	}
	panic("not implemented")
}

func newVariableParamExtractor(v st.Variable) func(base.Node) optional.Of[*data.Mutation] {
	return func(x base.Node) optional.Of[*data.Mutation] {
		return optional.Value(data.NewMutationWithVariableMapping(data.NewVariableMapping(v.Name(), x)))
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
