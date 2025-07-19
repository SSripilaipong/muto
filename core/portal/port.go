package portal

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type Port interface {
	Call(node base.Node) optional.Of[base.Node]
}
