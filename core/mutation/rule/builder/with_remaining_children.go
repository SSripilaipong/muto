package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
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
	if mutation == nil || mutation.RemainingParamChain().Size() == 0 {
		return optional.Value(node)
	}
	if !base.IsObjectNode(node) {
		node = base.NewOneLayerObject(node, []base.Node{})
	}
	return optional.Value[base.Node](base.UnsafeNodeToObject(node).ChainParams(mutation.RemainingParamChain()))
}
