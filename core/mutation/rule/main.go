package mutation

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/builder"
	ruleExtractor "github.com/SSripilaipong/muto/core/mutation/rule/extractor"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func New(rule st.Rule) func(t base.Object) optional.Of[base.Node] {
	return mutator.New(
		mutator.BuilderFunc(withRemainingChildren(builder.New(rule.Result()))),
		mutator.ExtractorFunc(ruleExtractor.New(rule.Pattern())),
	)
}

func withRemainingChildren(f func(*parameter.Parameter) optional.Of[base.Node]) func(*parameter.Parameter) optional.Of[base.Node] {
	return func(mutation *parameter.Parameter) optional.Of[base.Node] {
		node, ok := f(mutation).Return()
		if !ok {
			return optional.Empty[base.Node]()
		}
		if len(mutation.RemainingChildren()) == 0 {
			return optional.Value(node)
		}
		if base.IsClassNode(node) {
			return optional.Value[base.Node](base.NewObject(base.UnsafeNodeToClass(node), mutation.RemainingChildren()))
		}
		if base.IsObjectNode(node) {
			return optional.Value[base.Node](base.UnsafeNodeToObject(node).AppendChildren(mutation.RemainingChildren()))
		}
		return optional.Value[base.Node](base.NewObject(node, mutation.RemainingChildren()))
	}
}
