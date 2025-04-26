package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type normalMutationFunc struct {
	normal func(name string, obj base.Object) optional.Of[base.Node]
}

func newNormalMutationFunc(f func(name string, obj base.Object) optional.Of[base.Node]) normalMutationFunc {
	return normalMutationFunc{normal: f}
}

func (m normalMutationFunc) Active(_ string, _ base.Object) optional.Of[base.Node] {
	return optional.Empty[base.Node]()
}

func (m normalMutationFunc) Normal(name string, obj base.Object) optional.Of[base.Node] {
	return m.normal(name, obj)
}
