package time

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/optional"
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

func (t *sleepMutator) Mutate(_ base.Object) optional.Of[base.Node] {
	fmt.Println("sleeping")
	return optional.Value[base.Node](base.Null())
}

func (t *sleepMutator) VisitClass(mutator.ClassVisitor) {}
