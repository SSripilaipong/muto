package mutator

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
)

type Unit interface {
	Mutate(obj base.Object) optional.Of[base.Node]
	VisitClass(ClassVisitor)
}

type NamedUnit interface {
	Unit
	Name() string
}

func ToNamedUnit[T NamedUnit](x T) NamedUnit { return x }

type ClassVisitor interface {
	Visit(base.Class)
}

type ClassVisitorFunc func(base.Class)

func (f ClassVisitorFunc) Visit(c base.Class) { f(c) }
