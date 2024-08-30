package object

import (
	"phi-lang/common/fn"
	"phi-lang/common/slc"
	ruleMutation "phi-lang/core/mutation/rule"
	st "phi-lang/syntaxtree"
)

var NewMutatorsFromStatements = fn.Compose(reduceMutatorFromRules, mapFilterRuleFromStatement)

var reduceMutatorFromRules = slc.FoldGroup(mergeMutatorWithRule, st.RuleToPatternName)(NewMutator("", nil))

func mergeMutatorWithRule(t Mutator, rule st.Rule) Mutator {
	if t.name == "" {
		t.name = rule.PatternName()
	}
	t.mutationRules = append(t.mutationRules, ruleMutation.New(rule))
	return t
}

var mapFilterRuleFromStatement = fn.Compose(slc.Map(st.UnsafeStatementToRule), slc.Filter(st.IsRuleStatement))
