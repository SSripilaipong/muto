package mutation

import (
	"github.com/SSripilaipong/muto/core/mutation/rule/builder"
	ruleExtractor "github.com/SSripilaipong/muto/core/mutation/rule/extractor"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type RuleBuilder struct {
	builderFactory builder.SimplifiedNodeBuilderFactory
}

func NewRuleBuilder(classCollection builder.ClassCollection) RuleBuilder {
	return RuleBuilder{
		builderFactory: builder.NewSimplifiedNodeBuilderFactory(classCollection),
	}
}

func (b RuleBuilder) Build(rule st.Rule) mutator.NameWrapper {
	coreBuilder := b.NewResultBuilder(rule.Result())
	nodeBuilder := fixFreeObject(rule.Pattern(), rule.Result(), coreBuilder)
	return mutator.NewNameWrapper(
		rule.PatternName(),
		mutator.NewReconstructor(ruleExtractor.New(rule.Pattern()), nodeBuilder),
	)
}

func (b RuleBuilder) NewResultBuilder(rule stResult.SimplifiedNode) mutator.Builder {
	return b.builderFactory.NewBuilder(rule)
}
