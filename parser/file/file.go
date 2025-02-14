package file

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	psPattern "github.com/SSripilaipong/muto/parser/pattern"
	psResult "github.com/SSripilaipong/muto/parser/result"
	"github.com/SSripilaipong/muto/syntaxtree"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var file = ps.RsMap(base.NewFile, psBase.IgnoreLeadingLineBreak(psBase.IgnoreTrailingLineBreak(rsStatements)))

var Rule = ps.Map(rslt.Fmap(mergeRule), psBase.RsSpaceSeparated3(psPattern.RsNamedRule, psBase.RsEqualSign, psResult.RsNode))

var mergeRule = tuple.Fn3(func(p stPattern.NamedRule, _ psBase.Character, r stResult.Node) syntaxtree.Rule {
	return syntaxtree.NewRule(p, r)
})

var ActiveRule = ps.RsMap(mergeActiveRule, psBase.RsSpaceSeparated2(psBase.RsAtSign, Rule))

var mergeActiveRule = tuple.Fn2(func(_ psBase.Character, r syntaxtree.Rule) syntaxtree.ActiveRule {
	return syntaxtree.NewActiveRule(r.Pattern(), r.Result())
})
