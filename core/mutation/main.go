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

func globalMutationFromObjectMutators(ts []object.Mutator) func(base.ObjectLike) optional.Of[base.Node] {
	mutator := slc.ToMapValue(object.MutatorName)(append(ts, builtin.NewMutators()...))

	return func(obj base.ObjectLike) optional.Of[base.Node] {
		return mutator[obj.ClassName()].Mutate(obj)
	}
}
