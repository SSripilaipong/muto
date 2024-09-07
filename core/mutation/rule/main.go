package mutation

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/mutation/rule/builder"
	"muto/core/mutation/rule/data"
	"muto/core/mutation/rule/extractor"
	"muto/core/mutation/rule/mutator"
	st "muto/syntaxtree"
)

func New(rule st.Rule) func(t base.Object) optional.Of[base.Node] {
	return mutator.New(
		mutator.BuilderFunc(withRemainingChildren(builder.New(rule.Result()))),
		mutator.ExtractorFunc(extractor.New(rule.Pattern())),
	)
}

func withRemainingChildren(f func(*data.Mutation) optional.Of[base.Node]) func(*data.Mutation) optional.Of[base.Node] {
	return func(mutation *data.Mutation) optional.Of[base.Node] {
		node, ok := f(mutation).Return()
		if !ok {
			return optional.Empty[base.Node]()
		}
		if base.IsObjectNode(node) {
			node = base.UnsafeNodeToObject(node).AppendChildren(mutation.RemainingChildren())
		}
		return optional.Value(node)
	}
}
