package global

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

const tryMutatorName = "try"

type tryMutator struct {
}

func newTryMutator() *tryMutator {
	return &tryMutator{}
}

func (t *tryMutator) Name() string { return tryMutatorName }

func (t *tryMutator) Mutate(obj base.Object) optional.Of[base.Node] { // TODO check if it still works with param chain
	children := obj.ParamChain().DirectParams()
	if len(children) < 2 {
		return optional.Empty[base.Node]()
	}
	subject := base.NewCompoundObject(children[0], base.NewParamChain(slc.Pure(children[1:])))
	result := subject.Mutate()
	if result.IsEmpty() {
		return optional.Value[base.Node](base.EmptyTag)
	}
	return optional.Value[base.Node](base.NewCompoundObject(base.ValueTag, base.NewParamChain(slc.Pure([]base.Node{result.Value()}))))
}

func (t *tryMutator) VisitClass(mutator.ClassVisitor) {}

func objectStrictUnaryOp(f func(x base.Object) optional.Of[base.Node]) func(t base.Object) optional.Of[base.Node] {
	return strictUnaryOp(func(x base.Node) optional.Of[base.Node] {
		if !base.IsObjectNode(x) {
			return optional.Empty[base.Node]()
		}
		return f(base.UnsafeNodeToObject(x))
	})
}
