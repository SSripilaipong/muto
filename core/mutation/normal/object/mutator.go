package object

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type Mutator interface {
	Mutate(name string, obj base.Object) optional.Of[base.Node]
	SetGlobalMutator(_ Mutator)
	Name() string
}

type MutatorFunc func(name string, obj base.Object) optional.Of[base.Node]

func (m MutatorFunc) Mutate(name string, obj base.Object) optional.Of[base.Node] {
	return m(name, obj)
}

func (m MutatorFunc) SetGlobalMutator(_ Mutator) {}

func (m MutatorFunc) Name() string { return "" }

var _ Mutator = MutatorFunc(nil)

func MutatorName(t Mutator) string {
	return t.Name()
}

func ToMutator[T Mutator](x T) Mutator { return x }
