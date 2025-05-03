package mutation

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/slc"
	ruleMutation "github.com/SSripilaipong/muto/core/mutation/rule"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type ruleCollectionBuilder struct {
	builder ruleMutation.RuleBuilder
}

func newRuleCollectionBuilder(linker ClassLinker) ruleCollectionBuilder {
	return ruleCollectionBuilder{
		builder: ruleMutation.NewRuleBuilder(linker),
	}
}

func (b ruleCollectionBuilder) BuildRuleCollection(ss []stBase.Statement) mutator.RuleCollection {
	buildAll := slc.Map(fn.Compose(mutator.ToNamedObjectMutator, b.builder.Build))

	active := buildAll(st.FilterActiveRuleFromStatement(ss))
	normal := buildAll(st.FilterRuleFromStatement(ss))

	return mutator.NewRuleCollection(normal, active)
}

func (b ruleCollectionBuilder) NewRuleBuilder(rule st.Rule) mutator.NamedObjectMutator {
	return b.builder.Build(rule)
}

func (b ruleCollectionBuilder) NewResultBuilder(obj stResult.Object) mutator.Builder {
	return b.builder.NewResultBuilder(obj)
}
