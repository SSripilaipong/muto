package normal

import (
	"muto/common/fn"
	"muto/common/optional"
	"muto/common/slc"
	"muto/core/base"
	"muto/core/mutation/normal/builtin"
	"muto/core/mutation/normal/object"
)

var NewFromStatements = fn.Compose(newSelectiveMutator, object.NewMutatorsFromStatements)

func newSelectiveMutator(ms []object.Mutator) func(string, base.NamedObject) optional.Of[base.Node] {
	mutator := slc.ToMapValue(object.MutatorName)(append(ms, builtin.NewMutators()...))

	return func(name string, obj base.NamedObject) optional.Of[base.Node] {
		return mutator[name].Mutate(obj)
	}
}
