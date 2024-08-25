package extractor

import (
	"phi-lang/common/optional"
	"phi-lang/core/base"
	"phi-lang/core/mutation/rule/data"
	st "phi-lang/syntaxtree"
)

func newParamExtractors(r st.RulePattern) (result []func(base.Node) optional.Of[data.Mutation]) {
	for _, pattern := range r.Params() {
		result = append(result, newParamExtractor(pattern))
	}
	return
}

func newParamExtractor(p st.RuleParamPattern) func(base.Node) optional.Of[data.Mutation] {
	if st.IsRuleParamPatternVariable(p) {
		return newVariableParamExtractor(st.UnsafeRuleParamPatternToVariable(p))
	}
	panic("not implemented")
}

func newVariableParamExtractor(v st.Variable) func(base.Node) optional.Of[data.Mutation] {
	return func(x base.Node) optional.Of[data.Mutation] {
		return optional.Value(data.NewMutationWithVariableMapping(data.NewVariableMapping(v.Name(), x)))
	}
}
