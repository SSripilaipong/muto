package file

import (
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	psPattern "github.com/SSripilaipong/muto/parser/pattern"
	psResult "github.com/SSripilaipong/muto/parser/result"
	"github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var determinant = psPattern.Determinant()
var Rule = ps.Map(mergeRule, psBase.SpaceSeparated3(
	determinant,
	psBase.EqualSign,
	psResult.SimplifiedNodeInstant,
))

var mergeRule = tuple.Fn3(func(p stPattern.DeterminantObject, _ psBase.Character, r stResult.SimplifiedNode) syntaxtree.Rule {
	return syntaxtree.NewRule(p, r)
})

var ActiveRule = ps.Map(
	mergeActiveRule,
	psBase.SpaceSeparated2(psBase.AtSign, Rule),
)

var mergeActiveRule = tuple.Fn2(func(_ psBase.Character, r syntaxtree.Rule) syntaxtree.ActiveRule {
	return syntaxtree.NewActiveRule(r.Pattern(), r.Result())
})
