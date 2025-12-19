package base

import (
	"github.com/SSripilaipong/go-common/optional"
)

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
