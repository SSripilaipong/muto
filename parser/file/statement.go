package file

import (
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree"
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

var statements = ps.Map(
	aggregateStatements,
	ps.ToParser(psBase.IgnoreSpaceBetween2(
		statement.Legacy,
		ps.OptionalGreedyRepeat(ps.ToParser(psBase.WithLeadingLineBreak(statement.Legacy))).Legacy,
	)),
)

var aggregateStatements = tuple.Fn2(func(s base.Statement, ss []base.Statement) []base.Statement {
	return append([]base.Statement{s}, ss...)
})

var statement = ps.First(
	ps.Map(syntaxtree.ActiveRuleToStatement, ActiveRule),
	ps.Map(syntaxtree.RuleToStatement, Rule),
	ps.Map(syntaxtree.ImportToStatement, Command),
)
