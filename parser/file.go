package parser

import (
	ps "phi-lang/common/parsing"
	"phi-lang/common/tuple"
	st "phi-lang/parser/syntaxtree"
	"phi-lang/parser/tokenizer"
)

var file = ps.Map(st.NewFile, ignoreLeadingLineBreak(statements))

var statements = ps.OptionalGreedyRepeat(ignoreTrailingLineBreak(statement))
var statement = ps.Map(st.RuleToStatement, rule)

var rule = ps.Map(mergeRule, ps.Sequence3(rulePattern, equalSign, ruleResult))

var mergeRule = tuple.Fn3(func(p st.RulePattern, _ tokenizer.Token, r st.RuleResult) st.Rule {
	return st.NewRule(p, r)
})
