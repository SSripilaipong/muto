package mutator

import (
	"muto/common/optional"
	"muto/core/base"
)

type Builder[M any] interface {
	Build(M) optional.Of[base.Node]
}

func BuilderFunc[M any](f builderFunc[M]) Builder[M] {
	return f
}

type builderFunc[M any] func(M) optional.Of[base.Node]

func (f builderFunc[M]) Build(x M) optional.Of[base.Node] {
	return f(x)
}
