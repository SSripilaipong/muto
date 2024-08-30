package mutation

import (
	"phi-lang/common/fn"
	"phi-lang/common/optional"
	"phi-lang/common/slc"
	"phi-lang/core/base"
	"phi-lang/core/mutation/builtin"
	"phi-lang/core/mutation/object"
)

var NewFromStatements = fn.Compose(globalMutationFromObjectMutators, object.NewMutatorsFromStatements)

func globalMutationFromObjectMutators(ts []object.Mutator) func(base.Object) optional.Of[base.Node] {
	mutator := slc.ToMapValue(object.MutatorName)(append(ts, builtin.NewMutators()...))

	return func(obj base.Object) optional.Of[base.Node] {
		return mutator[obj.ClassName()].Mutate(obj)
	}
}
