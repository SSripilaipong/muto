package time

import (
	"time"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

const sleepMutatorName = "sleep"

type sleepMutator struct {
}

func newSleepMutator() *sleepMutator {
	return &sleepMutator{}
}

func (t *sleepMutator) Name() string { return sleepMutatorName }

func (t *sleepMutator) Mutate(obj base.Object) optional.Of[base.Node] {
	return base.StrictUnaryOp(func(x base.Node) optional.Of[base.Node] {
		if !base.IsNumberNode(x) {
			return optional.Empty[base.Node]()
		}
		num := base.UnsafeNodeToNumber(x)
		time.Sleep(time.Duration(num.Value().ToFloat() * 1e9))
		return optional.Value[base.Node](base.Null())
	})(obj.ParamChain())
}

func (t *sleepMutator) VisitClass(mutator.ClassVisitor) {}
