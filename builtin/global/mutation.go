package global

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type normalMutationFunc struct {
	normal func(obj base.Object) optional.Of[base.Node]
}

func newNormalMutationFunc(f func(obj base.Object) optional.Of[base.Node]) normalMutationFunc {
	return normalMutationFunc{normal: f}
}

func (m normalMutationFunc) Active(_ base.Object) optional.Of[base.Node] {
	return optional.Empty[base.Node]()
}

func (m normalMutationFunc) Normal(obj base.Object) optional.Of[base.Node] {
	return m.normal(obj)
}
