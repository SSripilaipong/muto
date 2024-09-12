package builtin

import (
	"muto/common/optional"
	"muto/core/base"
)

func terminate(t base.Object) optional.Of[base.Node] {
	return optional.Value[base.Node](t.ConfirmTermination())
}

func valueWithRemainingChildren(value base.Node, remainingChildren []base.Node) optional.Of[base.Node] {
	if len(remainingChildren) > 0 {
		newChildren := append([]base.Node{value}, remainingChildren...)
		return optional.Value(base.NamedObjectToNode(base.NewDataObject(newChildren)))
	}
	return optional.Value(value)
}