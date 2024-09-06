package mutation

import (
	"muto/common/fn"
	"muto/common/optional"
	"muto/common/slc"
	"muto/core/base"
	"muto/core/mutation/builtin"
	"muto/core/mutation/object"
)

var NewFromStatements = fn.Compose(globalMutationFromObjectMutators, object.NewMutatorsFromStatements)

func globalMutationFromObjectMutators(ts []object.Mutator) (recursiveMutate func(base.Object) optional.Of[base.Node]) {
	mutate := selectiveMutator(append(ts, builtin.NewMutators()...))

	return func(obj base.Object) optional.Of[base.Node] {
		return obj.MutateWithObjMutateFunc(mutate)
	}
}

func selectiveMutator(ms []object.Mutator) func(base.Object) optional.Of[base.Node] {
	mutator := slc.ToMapValue(object.MutatorName)(ms)

	return func(obj base.Object) optional.Of[base.Node] {
		return mutator[obj.ClassName()].Mutate(obj)
	}
}
