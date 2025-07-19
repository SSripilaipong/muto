package module

import (
	"github.com/SSripilaipong/muto/common/slc"
	ruleMutation "github.com/SSripilaipong/muto/core/mutation/rule"
	ruleMutationBuilder "github.com/SSripilaipong/muto/core/mutation/rule/builder"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
)

func BuildModuleFromStatements(ss []stBase.Statement) Uninitialized {
	ruleBuilder := ruleMutation.NewRuleBuilder(ruleMutationBuilder.NewSimplifiedNodeBuilderFactory())

	buildAll := slc.Map(ruleBuilder.BuildNamedUnit)
	active := buildAll(st.FilterActiveRuleFromStatement(ss))
	normal := buildAll(st.FilterRuleFromStatement(ss))

	collection := mutator.NewCollectionFromMutators(normal, active)
	return AsUninitialized(NewModule(collection, ruleBuilder))
}
