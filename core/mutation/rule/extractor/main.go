package extractor

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/mutation/rule/data"
	st "muto/syntaxtree"
)

func New(rule st.RulePattern) func(obj base.Object) optional.Of[*data.Mutation] {
	paramPart := rule.ParamPart()
	switch {
	case st.IsRulePatternParamPartTypeFixed(paramPart):
		return newForFixedParamPart(st.UnsafeRulePatternParamPartToArrayOfRuleParamPatterns(paramPart))
	case st.IsRulePatternParamPartTypeVariadic(paramPart):
		return newForVariadicParamPart(st.UnsafeRulePatternParamPartToVariadicParamPart(paramPart))
	}
	panic("not implemented")
}

func newForVariadicParamPart(paramPart st.RulePatternVariadicParamPart) func(obj base.Object) optional.Of[*data.Mutation] {
	panic("not implemented")
}

func newForFixedParamPart(params []st.RuleParamPattern) func(obj base.Object) optional.Of[*data.Mutation] {
	paramExtract := newParamExtractors(params)
	nConsumed := len(paramExtract)

	return func(obj base.Object) optional.Of[*data.Mutation] {
		if len(params) > len(obj.Children()) {
			return optional.Empty[*data.Mutation]()
		}

		mutation := data.NewMutation()
		for i, child := range obj.Children()[:nConsumed] {
			e := paramExtract[i](child)
			if e.IsEmpty() {
				return optional.Empty[*data.Mutation]()
			}
			m := mutation.Merge(e.Value())
			if m.IsEmpty() {
				return optional.Empty[*data.Mutation]()
			}
			mutation = m.Value()
		}
		return optional.Value(mutation.WithRemainingChildren(obj.Children()[nConsumed:]))
	}
}
