package base

import (
	"github.com/SSripilaipong/go-common/optional"
)

func StrictStructureUnaryOp(f func(s Structure) optional.Of[Node]) func(params ParamChain) optional.Of[Node] {
	return StrictUnaryOp(func(x Node) optional.Of[Node] {
		if !IsStructureNode(x) {
			return optional.Empty[Node]()
		}
		return f(UnsafeNodeToStructure(x))
	})
}

func StrictTagUnaryOp(f func(t Tag) optional.Of[Node]) func(params ParamChain) optional.Of[Node] {
	return StrictUnaryOp(func(x Node) optional.Of[Node] {
		if !IsTagNode(x) {
			return optional.Empty[Node]()
		}
		return f(UnsafeNodeToTag(x))
	})
}

func StrictObjectUnaryOp(f func(o Object) optional.Of[Node]) func(params ParamChain) optional.Of[Node] {
	return StrictUnaryOp(func(x Node) optional.Of[Node] {
		if !IsObjectNode(x) {
			return optional.Empty[Node]()
		}
		return f(UnsafeNodeToObject(x))
	})
}

func StrictNumberUnaryOp(f func(n Number) optional.Of[Node]) func(params ParamChain) optional.Of[Node] {
	return StrictUnaryOp(func(x Node) optional.Of[Node] {
		if !IsNumberNode(x) {
			return optional.Empty[Node]()
		}
		return f(UnsafeNodeToNumber(x))
	})
}

func StrictStringUnaryOp(f func(s String) optional.Of[Node]) func(params ParamChain) optional.Of[Node] {
	return StrictUnaryOp(func(x Node) optional.Of[Node] {
		if !IsStringNode(x) {
			return optional.Empty[Node]()
		}
		return f(UnsafeNodeToString(x))
	})
}

func StrictStringBinaryOp(f func(a, b String) optional.Of[Node]) func(params ParamChain) optional.Of[Node] {
	return func(params ParamChain) optional.Of[Node] {
		innerChildren := params.DirectParams()
		if len(innerChildren) != 2 {
			return optional.Empty[Node]()
		}
		if !IsStringNode(innerChildren[0]) || !IsStringNode(innerChildren[1]) {
			return optional.Empty[Node]()
		}

		remainingParams := params.SliceFromNodeOrEmpty(0, 2)
		return ProcessMutationResultWithParamChain(
			f(UnsafeNodeToString(innerChildren[0]), UnsafeNodeToString(innerChildren[1])),
			remainingParams,
		)
	}
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
