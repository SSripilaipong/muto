package repl

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	fileParser "github.com/SSripilaipong/muto/parser/file"
	psResult "github.com/SSripilaipong/muto/parser/result"
	"github.com/SSripilaipong/muto/syntaxtree"
	replSt "github.com/SSripilaipong/muto/syntaxtree/repl"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var ParseStatement = fn.Compose3(psBase.FilterResult, statement, psBase.StringToCharTokens)

var statement = ps.First(
	ps.Map(replSt.ToStatement, ps.ToParser(command)),
	ps.Map(mergeRule, fileParser.Rule),
	ps.Map(mergeActiveRule, fileParser.ActiveRule),
	ps.Map(mergeNode, ps.ToParser(psResult.SimplifiedNodeInstant)),
).Legacy

func mergeRule(r syntaxtree.Rule) replSt.Statement {
	return replSt.NewRule(r)
}

func mergeActiveRule(r syntaxtree.ActiveRule) replSt.Statement {
	return replSt.NewActiveRule(r)
}

func mergeNode(n stResult.SimplifiedNode) replSt.Statement {
	return replSt.NewNode(n)
}
