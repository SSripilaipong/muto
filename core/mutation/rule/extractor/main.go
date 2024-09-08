package extractor

import (
	"muto/common/fn"
	"muto/common/optional"
	"muto/core/base"
	"muto/core/mutation/rule/data"
	st "muto/syntaxtree"
)

func New(rule st.RulePattern) func(obj base.Object) optional.Of[*data.Mutation] {
	return newForParamPart(rule.ParamPart())
}

func newForParamPart(paramPart st.RulePatternParamPart) func(obj base.Object) optional.Of[*data.Mutation] {
	switch {
	case st.IsRulePatternParamPartTypeFixed(paramPart):
		return newForFixedParamPart(st.UnsafeRulePatternParamPartToArrayOfRuleParamPatterns(paramPart))
	case st.IsRulePatternParamPartTypeVariadic(paramPart):
		return newForVariadicParamPart(st.UnsafeRulePatternParamPartToVariadicParamPart(paramPart))
	}
	panic("not implemented")
}

func newForVariadicParamPart(paramPart st.RulePatternVariadicParamPart) func(obj base.Object) optional.Of[*data.Mutation] {
	switch {
	case st.IsRulePatternVariadicParamPartTypeRight(paramPart):
		return newForRightVariadicParamPart(st.UnsafeRulePatternVariadicParamPartTypeToRightVariadic(paramPart))
	case st.IsRulePatternVariadicParamPartTypeLeft(paramPart):
		return newForLeftVariadicParamPart(st.UnsafeRulePatternVariadicParamPartTypeToLeftVariadic(paramPart))
	}
	panic("not implemented")
}

func newForFixedParamPart(params []st.RuleParamPattern) func(obj base.Object) optional.Of[*data.Mutation] {
	return fn.Compose(extractChildrenNodes(params), base.ObjectToChildren)
}
