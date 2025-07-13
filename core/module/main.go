package module

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/slc"
	ruleMutation "github.com/SSripilaipong/muto/core/mutation/rule"
	ruleMutationBuilder "github.com/SSripilaipong/muto/core/mutation/rule/builder"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
)

func BuildModuleFromStatements(ss []stBase.Statement) Module {
	builder := ruleMutation.NewRuleBuilder(ruleMutationBuilder.NewSimplifiedNodeBuilderFactory())

	buildAll := slc.Map(fn.Compose(mutator.ToNamedObjectMutator, builder.Build))
	active := buildAll(st.FilterActiveRuleFromStatement(ss))
	normal := buildAll(st.FilterRuleFromStatement(ss))

	ruleCollection := mutator.NewRuleCollection(normal, active)
	return NewModule(ruleCollection, builder)
}
