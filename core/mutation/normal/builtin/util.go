package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

func valueWithRemainingChildren(value base.Node, remainingChildren []base.Node) optional.Of[base.Node] {
	if len(remainingChildren) > 0 {
		newChildren := append([]base.Node{value}, remainingChildren...)
		return optional.Value(base.ObjectToNode(base.NewDataObject(newChildren)))
	}
	return optional.Value(value)
}
