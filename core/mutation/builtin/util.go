package builtin

import (
	"phi-lang/common/optional"
	"phi-lang/core/base"
)

func terminate(t base.ObjectLike) optional.Of[base.Node] {
	return optional.Value[base.Node](t.Terminate())
}
