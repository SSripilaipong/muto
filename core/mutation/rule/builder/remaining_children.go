package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type chainRemainingChildren struct {
	wrapped mutator.Builder
}

func wrapChainRemainingChildren(wrapped mutator.Builder) chainRemainingChildren {
	return chainRemainingChildren{wrapped: wrapped}
}

func (x chainRemainingChildren) Build(mutation *parameter.Parameter) optional.Of[base.Node] {
	node, ok := x.wrapped.Build(mutation).Return()
	if !ok {
		return optional.Empty[base.Node]()
	}
	if mutation == nil || mutation.RemainingParamChain().Size() == 0 {
		return optional.Value(node)
	}

	if !base.IsObjectNode(node) {
		return optional.Value[base.Node](base.NewCompoundObject(node, mutation.RemainingParamChain()))
	}

	return optional.Value[base.Node](base.UnsafeNodeToObject(node).ChainParams(mutation.RemainingParamChain()))
}

type appendRemainingChildren struct {
	wrapped mutator.Builder
}

func wrapAppendRemainingChildren(wrapped mutator.Builder) appendRemainingChildren {
	return appendRemainingChildren{wrapped: wrapped}
}

func (x appendRemainingChildren) Build(mutation *parameter.Parameter) optional.Of[base.Node] {
	node, ok := x.wrapped.Build(mutation).Return()
	if !ok {
		return optional.Empty[base.Node]()
	}
	if mutation == nil {
		return optional.Value(node)
	}
	return appendRemainingParamToNode(mutation.RemainingParamChain())(node)
}

func appendRemainingParamToNode(paramChain base.ParamChain) func(base.Node) optional.Of[base.Node] {
	return func(node base.Node) optional.Of[base.Node] {
		if paramChain.Size() == 0 {
			return optional.Value(node)
		}
		if !base.IsObjectNode(node) {
			return optional.Value[base.Node](base.NewCompoundObject(node, paramChain))
		}

		return optional.Value[base.Node](base.UnsafeNodeToObject(node).AppendParams(paramChain))
	}
}
