package builtin

import (
	"github.com/SSripilaipong/muto/common/cliio"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

func NewMutators(cliReader CliReader, cliPrinter CliPrinter) []mutator.NamedObjectMutator {
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
		doMutator,
	}
}

func NewBuiltinMutatorsForStdio() []mutator.NamedObjectMutator { // TODO create builtin as a collection that's ready to be merged from here
	return NewMutators(
		CliReaderFunc(cliio.ReadInputOneLine),
		CliPrinterFunc(cliio.PrintStringWithNewLine),
	)
}
