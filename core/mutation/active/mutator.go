package active

import (
	"muto/common/fn"
	"muto/common/slc"
	"muto/core/mutation/normal/object"
	st "muto/syntaxtree"
)

var newMutatorsFromStatements = fn.Compose3(object.ReduceMutatorFromRules, slc.Map(activeRuleToRule), mapFilterActiveRuleFromStatement)

func activeRuleToRule(x st.ActiveRule) st.Rule {
	return st.NewRule(x.Pattern(), x.Result())
}

var mapFilterActiveRuleFromStatement = fn.Compose(slc.Map(st.UnsafeStatementToActiveRule), slc.Filter(st.IsActiveRuleStatement))
