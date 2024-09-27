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
	anonymousHead := ps.Or(
		ps.Map(booleanToRuleParamPattern, boolean),
		ps.Map(stringToRuleParamPattern, string_),
		ps.Map(numberToRuleParamPattern, number),
		inParentheses(nestedObjectRuleParamPattern_),
	)

	mergeNestedAnonymousObjectRuleParamPattern := tuple.Fn2(func(head st.RuleParamPattern, paramPart st.RulePatternParamPart) st.RuleParamPattern {
		return st.NewAnonymousRulePattern(head, paramPart)
	})

	return ps.Or(
		ps.Map(mergeNestedAnonymousObjectRuleParamPattern, ps.Sequence2(anonymousHead, rulePatternParamPart())),
		ps.Map(st.NamedRulePatternToRulePatternParam, namedRulePattern()),
		ps.Map(st.VariableRulePatternToRulePatternParam, variableRulePattern()),
	)(xs)
}
