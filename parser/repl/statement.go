package repl

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	fileParser "github.com/SSripilaipong/muto/parser/file"
	psResult "github.com/SSripilaipong/muto/parser/result"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	replSt "github.com/SSripilaipong/muto/syntaxtree/repl"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var ParseStatement = fn.Compose3(psBase.FilterStatement, statement, psBase.StringToCharTokens)

var statement = ps.Or(
	ps.RsMap(replSt.ToStatement, command),
	ps.RsMap(mergeRule, fileParser.Rule),
	ps.RsMap(mergeActiveRule, fileParser.ActiveRule),
	ps.RsMap(mergeNode, psResult.RsNode),
)

func mergeRule(r base.Rule) replSt.Statement {
	return replSt.NewRule(r)
}

func mergeActiveRule(r base.ActiveRule) replSt.Statement {
	return replSt.NewActiveRule(r)
}

func mergeNode(n stResult.Node) replSt.Statement {
	return replSt.NewNode(n)
}
