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
	mutatorCollection mutator.Collection
	builder           ruleMutation.RuleBuilder
}

func NewModule(mutatorCollection mutator.Collection, builder ruleMutation.RuleBuilder) Module {
	return Module{
		mutatorCollection: mutatorCollection,
		builder:           builder,
	}
}

func (p Module) GetClass(name string) *base.Class {
	m, exists := p.mutatorCollection.GetMutator(name).Return()
	if !exists {
		return base.NewUnlinkedClass(name)
	}
	return base.NewClass(name, m)
}

func (p Module) AppendNormal(mutator mutator.NamedUnit) {
	p.mutatorCollection.AppendNormal(mutator)
	mutator.LinkClass(p.mutatorCollection)
}

func (p Module) BuildRule(rule st.Rule) mutator.NamedUnit {
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

func (p Module) RuleCollection() mutator.Collection {
	return p.mutatorCollection
}
