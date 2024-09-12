package builtin

import (
	"muto/core/mutation/normal/object"
)

func NewMutators() []object.Mutator {
	return []object.Mutator{
		rollingAddMutator,
		rollingConcatMutator,
		stringMutator,
		addMutator,
		subtractMutator,
		concatMutator,
		isStringMutator,
		isNumberMutator,
		isBooleanMutator,
		multiplyMutator,
		divideMutator,
		equalMutator,
		greaterThanMutator,
		greaterThanOrEqualMutator,
		lessThanMutator,
		lessThanOrEqualMutator,
	}
}
