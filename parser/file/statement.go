package file

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

var ParseStatement = fn.Compose3(psBase.FilterStatement, statement, psBase.StringToCharTokens)

var rsStatements = ps.RsMap(aggregateStatements, psBase.RsIgnoreSpaceBetween2(statement, ps.RsOptionalGreedyRepeat(psBase.RsWithLeadingLineBreak(statement))))

var aggregateStatements = tuple.Fn2(func(s base.Statement, ss []base.Statement) []base.Statement {
	return append([]base.Statement{s}, ss...)
})

var statement = ps.RsFirst(
	ps.RsMap(base.ActiveRuleToStatement, ActiveRule),
	ps.RsMap(base.RuleToStatement, Rule),
)
