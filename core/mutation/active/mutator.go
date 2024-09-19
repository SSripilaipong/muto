package active

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/mutation/normal/object"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

var newMutatorsFromStatements = fn.Compose3(object.ReduceMutatorFromRules, slc.Map(activeRuleToRule), mapFilterActiveRuleFromStatement)

func activeRuleToRule(x st.ActiveRule) st.Rule {
	return st.NewRule(x.Pattern(), x.Result())
}

var mapFilterActiveRuleFromStatement = fn.Compose(slc.Map(st.UnsafeStatementToActiveRule), slc.Filter(st.IsActiveRuleStatement))
