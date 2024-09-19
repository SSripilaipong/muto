package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	"github.com/SSripilaipong/muto/parser/tokenizer"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

var file = ps.Map(st.NewFile, ignoreLeadingLineBreak(ignoreTrailingLineBreak(statements)))

var statements = ps.Map(aggregateStatements, ps.Sequence2(statement, ps.OptionalGreedyRepeat(withLeadingLineBreak(statement))))
var statement = ps.Or(
	ps.Map(mergeActiveRule, ps.Sequence2(atSign, rule)),
	ps.Map(st.RuleToStatement, rule),
)

var mergeActiveRule = tuple.Fn2(func(_ tokenizer.Token, r st.Rule) st.Statement {
	return st.NewActiveRule(r.Pattern(), r.Result())
})

var rule = ps.Map(mergeRule, ps.Sequence3(namedRulePattern(), equalSign, ruleResult))

var mergeRule = tuple.Fn3(func(p st.NamedRulePattern, _ tokenizer.Token, r st.RuleResult) st.Rule {
	return st.NewRule(p, r)
})

var aggregateStatements = tuple.Fn2(func(s st.Statement, ss []st.Statement) []st.Statement {
	return append([]st.Statement{s}, ss...)
})
