package mutator

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type Unit interface {
	Mutate(obj base.Object) optional.Of[base.Node]
	LinkClass(ClassLinker)
}

type NamedUnit interface {
	Unit
	Name() string
}

func ToNamedUnit[T NamedUnit](x T) NamedUnit { return x }

type ClassLinker interface {
	LinkClass(*base.Class)
}
