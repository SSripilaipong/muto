package mutation

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

var topLevelObjectFlatten = optional.Map(func(node base.Node) base.Node { // TODO remove this if not needed anymore
	if base.IsObjectNode(node) {
		obj := base.UnsafeNodeToObject(node)
		if len(obj.Children()) == 0 {
			return obj.Head()
		}
	}
	return node
})
