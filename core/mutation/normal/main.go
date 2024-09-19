package normal

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/normal/builtin"
	"github.com/SSripilaipong/muto/core/mutation/normal/object"
)

var NewFromStatements = fn.Compose(newSelectiveMutator, object.NewMutatorsFromStatements)

func newSelectiveMutator(ms []object.Mutator) func(string, base.NamedObject) optional.Of[base.Node] {
	mutator := slc.ToMapValue(object.MutatorName)(append(ms, builtin.NewMutators()...))

	return func(name string, obj base.NamedObject) optional.Of[base.Node] {
		return mutator[name].Mutate(obj)
	}
}
