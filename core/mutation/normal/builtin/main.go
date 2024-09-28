package builtin

import (
	"github.com/SSripilaipong/muto/core/mutation/normal/object"
)

func NewMutators() []object.Mutator {
	return []object.Mutator{
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
