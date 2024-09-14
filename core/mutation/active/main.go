package active

import (
	"muto/common/fn"
	"muto/common/optional"
	"muto/common/slc"
	"muto/core/base"
	"muto/core/mutation/normal/object"
)

var NewFromStatements = fn.Compose(newSelectiveMutator, newMutatorsFromStatements)

func newSelectiveMutator(ms []object.Mutator) func(string, base.NamedObject) optional.Of[base.Node] {
	mutator := slc.ToMapValue(object.MutatorName)(ms)

	return func(name string, obj base.NamedObject) optional.Of[base.Node] {
		return mutator[name].Mutate(obj)
	}
}
