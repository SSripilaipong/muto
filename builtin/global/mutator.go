package global

import (
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

func newForeignNormalMutators(cliReader CliReader, cliPrinter CliPrinter) []mutator.NamedObjectMutator {
	return []mutator.NamedObjectMutator{
		stringMutator,
		addMutator,
		subtractMutator,
		concatMutator,
		isStringMutator,
		isNumberMutator,
		isBooleanMutator,
		multiplyMutator,
		divideMutator,
		divIntegerMutator,
		modIntegerMutator,
		equalMutator,
		notEqualMutator,
		greaterThanMutator,
		greaterThanOrEqualMutator,
		lessThanMutator,
		lessThanOrEqualMutator,
		cliInputMutator(cliReader),
		cliPrintMutator(cliPrinter),
		andMutator,
		orMutator,
		notMutator,
		newTryMutator(),
	}
}
