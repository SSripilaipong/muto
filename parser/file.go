package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	psPattern "github.com/SSripilaipong/muto/parser/pattern"
	psResult "github.com/SSripilaipong/muto/parser/result"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var file = ps.RsMap(st.NewFile, psBase.IgnoreLeadingLineBreak(psBase.IgnoreTrailingLineBreak(rsStatements)))

var rsStatements = ps.RsMap(aggregateStatements, psBase.RsIgnoreSpaceBetween2(statement, ps.RsOptionalGreedyRepeat(psBase.RsWithLeadingLineBreak(statement))))

var statement = ps.Or(
	ps.RsMap(mergeActiveRule, psBase.RsSpaceSeparated2(psBase.RsAtSign, rsRule)),
	ps.RsMap(st.RuleToStatement, rsRule),
)

var mergeActiveRule = tuple.Fn2(func(_ psBase.Character, r st.Rule) st.Statement {
	return st.NewActiveRule(r.Pattern(), r.Result())
})

var rsRule = ps.Map(rslt.Fmap(mergeRule), psBase.RsSpaceSeparated3(psPattern.RsNamedRule, psBase.RsEqualSign, psResult.RsNode))

var mergeRule = tuple.Fn3(func(p stPattern.NamedRule, _ psBase.Character, r stResult.Node) st.Rule {
	return st.NewRule(p, r)
})

var aggregateStatements = tuple.Fn2(func(s st.Statement, ss []st.Statement) []st.Statement {
	return append([]st.Statement{s}, ss...)
})
