package mutation

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	"github.com/SSripilaipong/muto/syntaxtree"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Package struct {
	appendableRuleCollection mutator.RuleCollection
	linker                   ClassLinker
	builder                  ruleCollectionBuilder
}

func NewPackageFromStatements(ss []st.Statement, builtins []mutator.NamedObjectMutator) Package {
	linker := NewClassLinker()
	builder := newRuleCollectionBuilder(linker)

	linker.LinkCollection(mutator.NewRuleCollection(builtins, nil))

	ruleCollection := builder.BuildRuleCollection(ss)
	linker.LinkCollection(ruleCollection)

	return Package{
		appendableRuleCollection: ruleCollection,
		linker:                   linker,
		builder:                  builder,
	}
}

func (p Package) GetClass(name string) *base.Class {
	return p.linker.GetClass(name)
}

func (p Package) AppendNormalRule(mutator mutator.NamedObjectMutator) {
	r := p.appendableRuleCollection.AppendNormalRule(mutator)
	p.linker.Link(mutator.Name(), r)
}

func (p Package) BuildRule(rule syntaxtree.Rule) mutator.NamedObjectMutator {
	return p.builder.NewRuleBuilder(rule)
}

func (p Package) BuildNode(obj stResult.Object) optional.Of[base.Node] {
	return p.builder.NewResultBuilder(obj).Build(parameter.New())
}
