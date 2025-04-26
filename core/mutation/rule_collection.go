package mutation

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/slc"
	ruleMutation "github.com/SSripilaipong/muto/core/mutation/rule"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
)

type ruleCollectionBuilder struct {
	builder ruleMutation.RuleBuilder
}

func newRuleCollectionBuilder(classCollection *ClassCollectionAdapter) ruleCollectionBuilder {
	return ruleCollectionBuilder{
		builder: ruleMutation.NewRuleBuilder(classCollection),
	}
}

func (b ruleCollectionBuilder) Build(ss []stBase.Statement, builtins []mutator.NamedObjectMutator) mutator.RuleCollection {
	buildAll := slc.Map(fn.Compose(mutator.ToNamedObjectMutator, b.builder.Build))

	active := buildAll(st.FilterActiveRuleFromStatement(ss))
	normal := buildAll(st.FilterRuleFromStatement(ss))
	normalWithBuiltins := slc.CloneConcat(normal, builtins)

	return mutator.NewRuleCollection(normalWithBuiltins, active)
}
