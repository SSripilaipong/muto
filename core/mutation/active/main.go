package active

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/normal/object"
)

var NewFromStatements = fn.Compose(newSelectiveMutator, newMutatorsFromStatements)

func newSelectiveMutator(ms []object.Mutator) func(string, base.NamedObject) optional.Of[base.Node] {
	mutator := slc.ToMapValue(object.MutatorName)(ms)

	return func(name string, obj base.NamedObject) optional.Of[base.Node] {
		return mutator[name].Mutate(obj)
	}
}
