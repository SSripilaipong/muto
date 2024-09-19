package object

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/slc"
	ruleMutation "github.com/SSripilaipong/muto/core/mutation/rule"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

var NewMutatorsFromStatements = fn.Compose(ReduceMutatorFromRules, mapFilterRuleFromStatement)

var ReduceMutatorFromRules = slc.FoldGroup(mergeMutatorWithRule, st.RuleToPatternName)(NewMutator("", nil))

func mergeMutatorWithRule(t Mutator, rule st.Rule) Mutator {
	if t.name == "" {
		t.name = rule.PatternName()
	}
	t.mutationRules = append(t.mutationRules, ruleMutation.New(rule))
	return t
}

var mapFilterRuleFromStatement = fn.Compose(slc.Map(st.UnsafeStatementToRule), slc.Filter(st.IsRuleStatement))
