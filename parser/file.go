package parser

import (
	ps "muto/common/parsing"
	"muto/common/tuple"
	"muto/parser/tokenizer"
	st "muto/syntaxtree"
)

var file = ps.Map(st.NewFile, ignoreLeadingLineBreak(statements))

var statements = ps.OptionalGreedyRepeat(ignoreTrailingLineBreak(statement))
var statement = ps.Map(st.RuleToStatement, rule)

var rule = ps.Map(mergeRule, ps.Sequence3(namedRulePattern(), equalSign, ruleResult))

var mergeRule = tuple.Fn3(func(p st.NamedRulePattern, _ tokenizer.Token, r st.RuleResult) st.Rule {
	return st.NewRule(p, r)
})
