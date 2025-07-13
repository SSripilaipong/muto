package module

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	ruleMutation "github.com/SSripilaipong/muto/core/mutation/rule"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Module struct {
	ruleCollection mutator.RuleCollection
	builder        ruleMutation.RuleBuilder
}

func NewModule(ruleCollection mutator.RuleCollection, builder ruleMutation.RuleBuilder) Module {
	return Module{
		ruleCollection: ruleCollection,
		builder:        builder,
	}
}

func (p Module) GetClass(name string) *base.Class {
	m, exists := p.ruleCollection.GetMutator(name).Return()
	if !exists {
		return base.NewUnlinkedClass(name)
	}
	return base.NewClass(name, m)
}

func (p Module) AppendNormalRule(mutator mutator.NamedObjectMutator) {
	p.ruleCollection.AppendNormalRule(mutator)
	mutator.LinkClass(p.ruleCollection)
}

func (p Module) BuildRule(rule st.Rule) mutator.NamedObjectMutator {
	return p.builder.Build(rule)
}

func (p Module) BuildNode(node stResult.SimplifiedNode) optional.Of[base.Node] {
	builder := p.builder.NewResultBuilder(node)
	mutator.VisitClass(p.RuleCollection().LinkClass, builder)
	return builder.Build(parameter.New())
}

func (p Module) LoadGlobal(builtin Module) {
	p.RuleCollection().LoadGlobal(builtin.RuleCollection())
}

func (p Module) RuleCollection() mutator.RuleCollection {
	return p.ruleCollection
}
