package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	psPattern "github.com/SSripilaipong/muto/parser/pattern"
	psResult "github.com/SSripilaipong/muto/parser/result"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var file = ps.Map(st.NewFile, psBase.IgnoreLeadingLineBreak(psBase.IgnoreTrailingLineBreak(statements)))

var statements = ps.Map(aggregateStatements, psBase.IgnoreSpaceBetween2(statement, ps.OptionalGreedyRepeat(psBase.WithLeadingLineBreak(statement))))
var statement = ps.Or(
	ps.Map(mergeActiveRule, psBase.SpaceSeparated2(psBase.AtSign, rule)),
	ps.Map(st.RuleToStatement, rule),
)

var mergeActiveRule = tuple.Fn2(func(_ psBase.Character, r st.Rule) st.Statement {
	return st.NewActiveRule(r.Pattern(), r.Result())
})

var rule = ps.Map(mergeRule, psBase.SpaceSeparated3(psPattern.NamedRule(), psBase.EqualSign, psResult.Node()))

var mergeRule = tuple.Fn3(func(p stPattern.NamedRule, _ psBase.Character, r stResult.Node) st.Rule {
	return st.NewRule(p, r)
})

var aggregateStatements = tuple.Fn2(func(s st.Statement, ss []st.Statement) []st.Statement {
	return append([]st.Statement{s}, ss...)
})
