package mutation

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/mutation/normal/rule/builder"
	"muto/core/mutation/normal/rule/data"
	"muto/core/mutation/normal/rule/extractor"
	"muto/core/mutation/normal/rule/mutator"
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
		if len(mutation.RemainingChildren()) == 0 {
			return optional.Value(node)
		}
		if base.IsObjectNode(node) {
			return optional.Value[base.Node](base.UnsafeNodeToObject(node).AppendChildren(mutation.RemainingChildren()))
		}
		return optional.Value[base.Node](base.NewDataObject([]base.Node{node}).AppendChildren(mutation.RemainingChildren()))
	}
}
