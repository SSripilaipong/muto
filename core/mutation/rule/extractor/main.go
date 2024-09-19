package extractor

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/data"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func New(rule st.NamedRulePattern) func(obj base.Object) optional.Of[*data.Mutation] {
	return newForParamPart(rule.ParamPart(), nonStrictlyMatchChildren)
}

func newForParamPart(paramPart st.RulePatternParamPart, nChildrenMatch func(nP int, nC int) bool) func(obj base.Object) optional.Of[*data.Mutation] {
	switch {
	case st.IsRulePatternParamPartTypeFixed(paramPart):
		return newForFixedParamPart(st.UnsafeRulePatternParamPartToArrayOfRuleParamPatterns(paramPart), nChildrenMatch)
	case st.IsRulePatternParamPartTypeVariadic(paramPart):
		return newForVariadicParamPart(st.UnsafeRulePatternParamPartToVariadicParamPart(paramPart), nChildrenMatch)
	}
	panic("not implemented")
}

func newForVariadicParamPart(paramPart st.RulePatternVariadicParamPart, nChildrenMatch func(nP int, nC int) bool) func(obj base.Object) optional.Of[*data.Mutation] {
	switch {
	case st.IsRulePatternVariadicParamPartTypeRight(paramPart):
		return newForRightVariadicParamPart(st.UnsafeRulePatternVariadicParamPartTypeToRightVariadic(paramPart), nChildrenMatch)
	case st.IsRulePatternVariadicParamPartTypeLeft(paramPart):
		return newForLeftVariadicParamPart(st.UnsafeRulePatternVariadicParamPartTypeToLeftVariadic(paramPart), nChildrenMatch)
	}
	panic("not implemented")
}

func newForFixedParamPart(params []st.RuleParamPattern, nChildrenMatch func(nP int, nC int) bool) func(obj base.Object) optional.Of[*data.Mutation] {
	return fn.Compose(extractChildrenNodes(params, nChildrenMatch), base.ObjectToChildren)
}
