package extractor

import (
	"phi-lang/common/fn"
	"phi-lang/common/optional"
	"phi-lang/common/slc"
	"phi-lang/core/base"
	"phi-lang/core/mutation/rule/data"
	st "phi-lang/syntaxtree"
)

func newParamExtractors(x st.RulePattern) []func(base.Node) optional.Of[*data.Mutation] {
	return fn.Compose(slc.Map(newParamExtractor), st.ParamsOfRulePattern)(x)
}

func newParamExtractor(p st.RuleParamPattern) func(base.Node) optional.Of[*data.Mutation] {
	if st.IsRuleParamPatternVariable(p) {
		return newVariableParamExtractor(st.UnsafeRuleParamPatternToVariable(p))
	} else if st.IsRuleParamPatternString(p) {
		return newStringParamExtractor(st.UnsafeRuleParamPatternToString(p))
	} else if st.IsRuleParamPatternNumber(p) {
		return newNumberParamExtractor(st.UnsafeRuleParamPatternToNumber(p))
	} else if st.IsRuleParamPatternNestedRulePattern(p) {
		return newNestedRuleExtractor(st.UnsafeRuleParamPatternToRulePattern(p))
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

func newNestedRuleExtractor(p st.RulePattern) func(base.Node) optional.Of[*data.Mutation] {
	extract := New(p)
	return func(x base.Node) optional.Of[*data.Mutation] {
		if base.IsObjectNode(x) {
			return extract(UnsafeNodeToObjectLike(x))
		}
		return optional.Empty[*data.Mutation]()
	}
}
