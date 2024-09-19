package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func nestedObjectRuleParamPattern(xs []tk.Token) []tuple.Of2[st.RuleParamPattern, []tk.Token] {
	return inParentheses(nestedObjectRuleParamPattern_)(xs)
}

func nestedObjectRuleParamPattern_(xs []tk.Token) []tuple.Of2[st.RuleParamPattern, []tk.Token] {
	return ps.Or(
		ps.Map(mergeNestedAnonymousObjectRuleParamPattern, ps.Sequence2(inParentheses(nestedObjectRuleParamPattern_), rulePatternParamPart())),
		ps.Map(st.NamedRulePatternToRulePatternParam, namedRulePattern()),
		ps.Map(st.VariableRulePatternToRulePatternParam, variableRulePattern()),
	)(xs)
}

var mergeNestedAnonymousObjectRuleParamPattern = tuple.Fn2(func(head st.RuleParamPattern, paramPart st.RulePatternParamPart) st.RuleParamPattern {
	return st.NewAnonymousRulePattern(head, paramPart)
})

func inParentheses[T any](x func([]tk.Token) []tuple.Of2[T, []tk.Token]) func([]tk.Token) []tuple.Of2[T, []tk.Token] {
	return ps.Map(withoutParenthesis[T], ps.Sequence3(openParenthesis, x, closeParenthesis))
}

func withoutParenthesis[T any](x tuple.Of3[tk.Token, T, tk.Token]) T {
	return x.X2()
}
