package base

import (
	"github.com/SSripilaipong/muto/common/optional"
)

type NameWiseMutation interface {
	Active(name string, obj Object) optional.Of[Node]
	Normal(name string, obj Object) optional.Of[Node]
}

func StrictUnaryOp(f func(x Node) optional.Of[Node]) func(params ParamChain) optional.Of[Node] {
	return func(params ParamChain) optional.Of[Node] {
		innerChildren := params.DirectParams()
		if len(innerChildren) != 1 {
			return optional.Empty[Node]()
		}

		remainingParams := params.SliceFromNodeOrEmpty(0, 1)
		return ProcessMutationResultWithParamChain(f(innerChildren[0]), remainingParams)
	}
}

func ProcessMutationResultWithParamChain(r optional.Of[Node], remainingChain ParamChain) optional.Of[Node] {
	result, ok := r.Return()
	if !ok {
		return optional.Empty[Node]()
	}

	switch {
	case IsObjectNode(result):
		obj := UnsafeNodeToObject(result)
		return optional.Value[Node](obj.ChainParams(remainingChain))
	default:
		if remainingChain.Size() == 0 {
			return optional.Value(result)
		}
		return optional.Value[Node](NewCompoundObject(result, remainingChain))
	}
}
