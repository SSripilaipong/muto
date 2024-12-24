package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type withRemainingChildren struct {
	wrapped mutator.Builder
}

func wrapWithRemainingChildren(wrapped mutator.Builder) withRemainingChildren {
	return withRemainingChildren{wrapped: wrapped}
}

func (x withRemainingChildren) Build(mutation *parameter.Parameter) optional.Of[base.Node] {
	node, ok := x.wrapped.Build(mutation).Return()
	if !ok {
		return optional.Empty[base.Node]()
	}
	if mutation == nil || len(mutation.RemainingChildren()) == 0 {
		return optional.Value(node)
	}
	if base.IsClassNode(node) {
		return optional.Value[base.Node](base.NewObject(base.UnsafeNodeToClass(node), base.NewParamChain(slc.Pure(mutation.RemainingChildren()))))
	}
	if base.IsObjectNode(node) {
		return optional.Value[base.Node](base.UnsafeNodeToObject(node).AppendChildren(mutation.RemainingChildren()))
	}
	return optional.Value[base.Node](base.NewObject(node, base.NewParamChain(slc.Pure(mutation.RemainingChildren()))))
}
