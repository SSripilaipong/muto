package mutator

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type ObjectMutator interface {
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

type GlobalMutatorAware interface {
	SetGlobalMutator(_ NameBasedMutator)
}

type NamedObjectMutator interface {
	Named
	ObjectMutator
}

type NamedAppendableObjectMutator interface {
	NamedObjectMutator
	Append(r NamedObjectMutator)
}
