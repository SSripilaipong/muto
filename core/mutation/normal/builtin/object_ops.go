package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

const tryMutatorName = "try"

type tryMutator struct {
	globalMutator mutator.NameBasedMutator
}

func newTryMutator() *tryMutator {
	return &tryMutator{
		globalMutator: nil, // assigned later
	}
}

func (t *tryMutator) Name() string { return tryMutatorName }

func (t *tryMutator) MutateByName(name string, obj base.Object) optional.Of[base.Node] {
	if name != tryMutatorName {
		return optional.Empty[base.Node]()
	}
	return t.Mutate(obj)
}

func (t *tryMutator) Mutate(obj base.Object) optional.Of[base.Node] {
	children := obj.Children()
	if len(children) < 2 {
		return optional.Empty[base.Node]()
	}
	subject := base.NewObject(children[0], children[1:])
	if bb, ok := subject.BubbleUp().Return(); ok {
		if base.IsObjectNode(bb) {
			subject = base.UnsafeNodeToObject(bb)
		} else {
			panic("unexpected error")
		}
	}
	result := subject.Mutate(newNormalMutationFunc(t.globalMutator.MutateByName))
	if result.IsEmpty() {
		return optional.Value[base.Node](base.NewObject(base.EmptyTag, []base.Node{}))
	}
	return optional.Value[base.Node](base.NewObject(base.ValueTag, []base.Node{result.Value()}))
}

func (t *tryMutator) SetGlobalMutator(m mutator.NameBasedMutator) {
	t.globalMutator = m
}

var _ mutator.NameBounded = (*tryMutator)(nil)
