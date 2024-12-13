package active

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/slc"
	ruleMutation "github.com/SSripilaipong/muto/core/mutation/rule"
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

var newMutatorsFromStatements = fn.Compose3(slc.Map(ruleMutation.New), slc.Map(activeRuleToRule), mapFilterActiveRuleFromStatement)

func activeRuleToRule(x base.ActiveRule) base.Rule {
	return base.NewRule(x.Pattern(), x.Result())
}

var mapFilterActiveRuleFromStatement = fn.Compose(slc.Map(base.UnsafeStatementToActiveRule), slc.Filter(base.IsActiveRuleStatement))
