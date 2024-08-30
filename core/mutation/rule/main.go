package mutation

import (
	"phi-lang/common/optional"
	"phi-lang/core/base"
	"phi-lang/core/mutation/rule/builder"
	"phi-lang/core/mutation/rule/extractor"
	"phi-lang/core/mutation/rule/mutator"
	st "phi-lang/syntaxtree"
)

func New(rule st.Rule) func(t base.ObjectLike) optional.Of[base.Node] {
	return mutator.New(
		mutator.BuilderFunc(builder.New(rule.Result())),
		mutator.ExtractorFunc(extractor.New(rule.Pattern())),
	)
}
