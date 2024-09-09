package parser

import (
	ps "muto/common/parsing"
	"muto/common/tuple"
	tk "muto/parser/tokenizer"
	st "muto/syntaxtree"
)

func nestedObjectRuleParamPattern(xs []tk.Token) []tuple.Of2[st.RuleParamPattern, []tk.Token] {
	obj := ps.Or(
		ps.Map(st.NamedRulePatternToRulePatternParam, namedRulePattern()),
		ps.Map(st.VariableRulePatternToRulePatternParam, variableRulePattern()),
	)
	return ps.Map(withoutParenthesis, ps.Sequence3(openParenthesis, obj, closeParenthesis))(xs)
}

func withoutParenthesis[T any](x tuple.Of3[tk.Token, T, tk.Token]) T {
	return x.X2()
}
