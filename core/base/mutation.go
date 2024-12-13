package base

import "github.com/SSripilaipong/muto/common/optional"

type NameWiseMutation interface {
	Active(name string, obj Object) optional.Of[Node]
	Normal(name string, obj Object) optional.Of[Node]
}

func unaryOp(f func(x Node) optional.Of[Node]) func(children []Node) optional.Of[Node] {
	return func(children []Node) optional.Of[Node] {
		if len(children) == 0 {
			return optional.Empty[Node]()
		}
		return ProcessMutationResultWithChildren(f(children[0]), children[1:])
	}
}

func ProcessMutationResultWithChildren(r optional.Of[Node], otherChildren []Node) optional.Of[Node] {
	result, ok := r.Return()
	if !ok {
		return optional.Empty[Node]()
	}

	switch {
	case IsObjectNode(result):
		obj := UnsafeNodeToObject(result)
		return optional.Value[Node](obj.AppendChildren(otherChildren))
	default:
		if len(otherChildren) == 0 {
			return optional.Value(result)
		}
		return optional.Value[Node](NewObject(result, otherChildren))
	}
}
