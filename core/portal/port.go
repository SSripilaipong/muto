package portal

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
)

type Port interface {
	Call(node base.Node) optional.Of[base.Node]
}
