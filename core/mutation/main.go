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

func globalMutationFromObjectMutators(ts []object.Mutator) (recursiveMutate func(base.ObjectLike) optional.Of[base.Node]) {
	mutate := selectiveMutator(append(ts, builtin.NewMutators()...))

	recursiveMutate = func(obj base.ObjectLike) optional.Of[base.Node] {
		if obj.IsTerminationConfirmed() {
			return optional.Empty[base.Node]()
		}
		for i, child := range obj.Children() {
			if !child.IsTerminationConfirmed() {
				childObj := base.UnsafeNodeToObject(child)
				if newChild := recursiveMutate(childObj); newChild.IsNotEmpty() {
					return optional.Value[base.Node](obj.ReplaceChild(i, newChild.Value()))
				}
				obj = obj.ReplaceChild(i, childObj.ConfirmTermination())
			}
		}
		return mutate(obj)
	}
	return
}

func selectiveMutator(ms []object.Mutator) func(base.ObjectLike) optional.Of[base.Node] {
	mutator := slc.ToMapValue(object.MutatorName)(ms)

	return func(obj base.ObjectLike) optional.Of[base.Node] {
		return mutator[obj.ClassName()].Mutate(obj)
	}
}
