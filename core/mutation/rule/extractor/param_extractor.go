package extractor

import (
	"phi-lang/common/fn"
	"phi-lang/common/optional"
	"phi-lang/common/slc"
	"phi-lang/core/base"
	"phi-lang/core/mutation/rule/data"
	st "phi-lang/syntaxtree"
)

var newParamExtractors = fn.Compose(slc.Map(newParamExtractor), st.ParamsOfRulePattern)

func newParamExtractor(p st.RuleParamPattern) func(base.Node) optional.Of[*data.Mutation] {
	if st.IsRuleParamPatternVariable(p) {
		return newVariableParamExtractor(st.UnsafeRuleParamPatternToVariable(p))
	} else if st.IsRuleParamPatternString(p) {
		return newStringParamExtractor(st.UnsafeRuleParamPatternToString(p))
	}
	panic("not implemented")
}

func newVariableParamExtractor(v st.Variable) func(base.Node) optional.Of[*data.Mutation] {
	return func(x base.Node) optional.Of[*data.Mutation] {
		return optional.Value(data.NewMutationWithVariableMapping(data.NewVariableMapping(v.Name(), x)))
	}
}

func newStringParamExtractor(v st.String) func(base.Node) optional.Of[*data.Mutation] {
	return func(x base.Node) optional.Of[*data.Mutation] {
		if base.IsStringNode(x) && base.UnsafeNodeToString(x).Value() == v.StringValue() {
			return optional.Value(data.NewMutation())
		}
		return optional.Empty[*data.Mutation]()
	}
}
