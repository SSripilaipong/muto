package file

import (
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree"
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

var rsStatements = ps.RsMap(aggregateStatements, psBase.RsIgnoreSpaceBetween2(statement, ps.RsOptionalGreedyRepeat(psBase.RsWithLeadingLineBreak(statement))))

var aggregateStatements = tuple.Fn2(func(s base.Statement, ss []base.Statement) []base.Statement {
	return append([]base.Statement{s}, ss...)
})

var statement = ps.RsFirst(
	ps.RsMap(syntaxtree.ActiveRuleToStatement, ActiveRule),
	ps.RsMap(syntaxtree.RuleToStatement, Rule),
	ps.RsMap(syntaxtree.ImportToStatement, Command),
)
