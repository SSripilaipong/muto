package module

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	ruleMutation "github.com/SSripilaipong/muto/core/mutation/rule"
	ruleMutationBuilder "github.com/SSripilaipong/muto/core/mutation/rule/builder"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Module struct {
	ruleCollection mutator.RuleCollection
	linker         ClassLinker
	builder        ruleMutation.RuleBuilder
}

func BuildModuleFromStatements(ss []stBase.Statement, builtin Module) Module {
	linker := builtin.ClonedLinker()
	builder := ruleMutation.NewRuleBuilder(ruleMutationBuilder.NewSimplifiedNodeBuilderFactory(linker))

	buildAll := slc.Map(fn.Compose(mutator.ToNamedObjectMutator, builder.Build))
	active := buildAll(st.FilterActiveRuleFromStatement(ss))
	normal := buildAll(st.FilterRuleFromStatement(ss))

	ruleCollection := mutator.NewRuleCollection(normal, active)
	linker.LinkCollection(ruleCollection)

	return Module{
		ruleCollection: ruleCollection,
		linker:         linker,
		builder:        builder,
	}
}

func NewModule(ruleCollection mutator.RuleCollection, linker ClassLinker, builder ruleMutation.RuleBuilder) Module {
	return Module{
		ruleCollection: ruleCollection,
		linker:         linker,
		builder:        builder,
	}
}

func (p Module) GetOrCreateClass(name string) *base.Class {
	return p.linker.GetOrCreateClass(name)
}

func (p Module) AppendNormalRule(mutator mutator.NamedObjectMutator) {
	r := p.ruleCollection.AppendNormalRule(mutator)
	p.linker.Link(mutator.Name(), r)
}

func (p Module) BuildRule(rule st.Rule) mutator.NamedObjectMutator {
	return p.builder.Build(rule)
}

func (p Module) BuildNode(obj stResult.Object) optional.Of[base.Node] {
	return p.builder.NewResultBuilder(obj).Build(parameter.New())
}

func (p Module) ClonedLinker() ClassLinker {
	return p.linker.Clone()
}
