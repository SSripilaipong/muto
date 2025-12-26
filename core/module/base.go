package module

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	ruleMutation "github.com/SSripilaipong/muto/core/mutation/rule"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	"github.com/SSripilaipong/muto/core/portal"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Base struct {
	mutatorCollection mutator.Collection
	builder           ruleMutation.RuleBuilder
}

func NewBase(mutatorCollection mutator.Collection, builder ruleMutation.RuleBuilder) Base {
	return Base{
		mutatorCollection: mutatorCollection,
		builder:           builder,
	}
}

func (p Base) GetClass(name string) base.Class {
	m, exists := p.MutatorCollection().GetMutator(name).Return()
	if !exists {
		return base.NewUnlinkedRuleBasedClass(name)
	}
	return base.NewRuleBasedClass(name, m)
}

func (p Base) AppendNormal(m mutator.NamedUnit) {
	p.MutatorCollection().AppendNormal(m)
}

func (p Base) ExtendCollection(collection mutator.Collection) {
	p.MutatorCollection().Extend(collection)
	p.MutatorCollection().VisitClass(mutator.ClassVisitorFunc(collection.LaxLinkClass))
}

func (p Base) ExtendImportedCollection(moduleName string, collection mutator.Collection) {
	p.MutatorCollection().Extend(collection)
	p.MutatorCollection().MapImportedCollection(moduleName, collection)
}

func (p Base) BuildRule(rule st.Rule) mutator.NamedUnit {
	return p.builder.Build(rule)
}

func (p Base) BuildNode(node stResult.SimplifiedNode) optional.Of[base.Node] {
	builder := p.builder.NewResultBuilder(node)
	mutator.VisitClass(p.MutatorCollection().LinkClass, builder)
	return builder.Build(parameter.New())
}

func (p Base) MountPortal(q *portal.Portal) {
	p.MutatorCollection().MountPortal(q)
}

func (p Base) MapImportedModules(mapping ImportMapping) {
	p.MutatorCollection().MapImportedCollectionMapping(mapping)
}

func (p Base) MutatorCollection() mutator.Collection {
	return p.mutatorCollection
}
