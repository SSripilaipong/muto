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

func newSelectiveMutator(ms []object.Mutator) func(string, base.Object) optional.Of[base.Node] {
	ms = append(ms, builtin.NewMutators()...)
	mutator := slc.ToMapValue(object.MutatorName)(ms)

	f := func(name string, obj base.Object) optional.Of[base.Node] {
		if m, ok := mutator[name]; ok {
			return m.Mutate(name, obj)
		}
		return optional.Empty[base.Node]()
	}
	for _, m := range ms {
		m.SetGlobalMutator(object.MutatorFunc(f))
	}
	return f
}
