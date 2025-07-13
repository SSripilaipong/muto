package mutator

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type RuleMutator interface {
	Active(obj base.Object) optional.Of[base.Node]
	Normal(obj base.Object) optional.Of[base.Node]
}

type ObjectMutator interface {
	LinkClass(ClassLinker)
	Mutate(obj base.Object) optional.Of[base.Node]
}

type NameBasedMutator interface {
	MutateByName(name string, obj base.Object) optional.Of[base.Node]
}

type Named interface {
	Name() string
}

type NameBounded interface {
	NameBasedMutator
	Named
}

type NameBoundedFunc func(name string, obj base.Object) optional.Of[base.Node]

func (m NameBoundedFunc) MutateByName(name string, obj base.Object) optional.Of[base.Node] {
	return m(name, obj)
}

func (m NameBoundedFunc) Name() string { return "" }

var _ NameBounded = NameBoundedFunc(nil)

type NamedObjectMutator interface {
	Named
	ObjectMutator
}

func ToNamedObjectMutator[T NamedObjectMutator](x T) NamedObjectMutator { return x }

type NamedRuleMutator interface {
	Named
	RuleMutator
}

type ClassLinker interface {
	LinkClass(*base.Class)
}
