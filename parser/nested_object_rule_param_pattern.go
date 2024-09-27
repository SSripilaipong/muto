package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func nestedObjectRuleParamPattern(xs []tk.Token) []tuple.Of2[stPattern.Param, []tk.Token] {
	return inParentheses(nestedObjectRuleParamPattern_)(xs)
}

func nestedObjectRuleParamPattern_(xs []tk.Token) []tuple.Of2[stPattern.Param, []tk.Token] {
	anonymousHead := ps.Or(
		ps.Map(booleanToRuleParamPattern, boolean),
		ps.Map(stringToRuleParamPattern, string_),
		ps.Map(numberToRuleParamPattern, number),
		inParentheses(nestedObjectRuleParamPattern_),
	)

	mergeNestedAnonymousObjectRuleParamPattern := tuple.Fn2(func(head stPattern.Param, paramPart stPattern.ParamPart) stPattern.Param {
		return stPattern.NewAnonymousRule(head, paramPart)
	})

	return ps.Or(
		ps.Map(mergeNestedAnonymousObjectRuleParamPattern, ps.Sequence2(anonymousHead, rulePatternParamPart())),
		ps.Map(stPattern.NamedRuleToParam, namedRulePattern()),
		ps.Map(stPattern.VariableRulePatternToRulePatternParam, variableRulePattern()),
	)(xs)
}
