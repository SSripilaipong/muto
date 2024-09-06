package mutation

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/mutation/rule/builder"
	"muto/core/mutation/rule/extractor"
	"muto/core/mutation/rule/mutator"
	st "muto/syntaxtree"
)

func New(rule st.Rule) func(t base.Object) optional.Of[base.Node] {
	return mutator.New(
		mutator.BuilderFunc(builder.New(rule.Result())),
		mutator.ExtractorFunc(extractor.New(rule.Pattern())),
	)
}
