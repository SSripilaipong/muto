package mutator

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type NameWrapper struct {
	ObjectMutator
	name string
}

func NewNameWrapper(name string, mutator ObjectMutator) NameWrapper {
	return NameWrapper{ObjectMutator: mutator, name: name}
}

func (n NameWrapper) MutateByName(name string, obj base.Object) optional.Of[base.Node] {
	if name != n.name {
		return optional.Empty[base.Node]()
	}
	return n.Mutate(obj)
}

func (n NameWrapper) Name() string {
	return n.name
}

var _ NameBounded = NameWrapper{}
