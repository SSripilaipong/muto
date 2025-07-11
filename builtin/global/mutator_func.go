package global

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

type objectMutatorFunc func(t base.Object) optional.Of[base.Node]

func (f objectMutatorFunc) Mutate(obj base.Object) optional.Of[base.Node] {
	return f(obj)
}

func NewRuleBasedMutatorFromFunctions(name string, mutationRules []func(t base.Object) optional.Of[base.Node]) mutator.ClassMutator {
	mutators := slc.Map(func(f func(t base.Object) optional.Of[base.Node]) mutator.ObjectMutator {
		return objectMutatorFunc(f)
	})(mutationRules)
	return mutator.NewClassMutator(name, mutators)
}
