package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

func unaryOp(f func(x base.Node) optional.Of[base.Node]) func(t base.Object) optional.Of[base.Node] {
	return func(t base.Object) optional.Of[base.Node] {
		params := t.ParamChain()
		innerChildren := params.DirectParams()
		if len(innerChildren) < 1 {
			return optional.Empty[base.Node]()
		}

		remainingParams := params.SliceFromNodeOrEmpty(0, 1)
		return base.ProcessMutationResultWithParamChain(f(innerChildren[0]), remainingParams)
	}
}

func binaryOp(f func(x, y base.Node) optional.Of[base.Node]) func(t base.Object) optional.Of[base.Node] {
	return func(t base.Object) optional.Of[base.Node] {
		params := t.ParamChain()
		innerChildren := params.DirectParams()
		if len(innerChildren) < 2 {
			return optional.Empty[base.Node]()
		}

		remainingParams := params.SliceFromNodeOrEmpty(0, 2)
		return base.ProcessMutationResultWithParamChain(f(innerChildren[0], innerChildren[1]), remainingParams)
	}
}

func leftVariadicUnaryOp(f func(xs []base.Node, x base.Node) optional.Of[base.Node]) func(t base.Object) optional.Of[base.Node] {
	return func(t base.Object) optional.Of[base.Node] {
		params := t.ParamChain()
		innerChildren := params.DirectParams()
		length := len(innerChildren)
		if length < 1 {
			return optional.Empty[base.Node]()
		}

		remainingParams := base.NewParamChain(nil).AppendAll(params.SliceFromOrEmpty(1))
		return base.ProcessMutationResultWithParamChain(f(innerChildren[:length-1], innerChildren[length-1]), remainingParams)
	}
}
