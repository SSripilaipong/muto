package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

type CliReader interface {
	Read() rslt.Of[string]
}

type CliReaderFunc func() rslt.Of[string]

func (r CliReaderFunc) Read() rslt.Of[string] { return r() }

type CliPrinter interface {
	Print(string)
}

type CliPrinterFunc func(string)

func (p CliPrinterFunc) Print(s string) { p(s) }

func cliPrintMutator(printer CliPrinter) mutator.ClassMutator {
	return NewRuleBasedMutatorFromFunctions("print!", slc.Pure(strictUnaryOp(func(x base.Node) optional.Of[base.Node] {
		if base.IsStringNode(x) {
			printer.Print(base.UnsafeNodeToString(x).Value())
			return optional.Value[base.Node](base.NewClass("$"))
		}
		return optional.Empty[base.Node]()
	})))
}

func cliInputMutator(reader CliReader) mutator.ClassMutator {
	return NewRuleBasedMutatorFromFunctions("input!", slc.Pure(strictUnaryOp(func(x base.Node) optional.Of[base.Node] {
		if base.IsClassNode(x) && base.UnsafeNodeToClass(x).Name() == "$" {
			s, err := reader.Read().Return()
			if err != nil {
				return optional.Value[base.Node](base.NewErrorWithMessage(err.Error()))
			}
			return optional.Value[base.Node](base.NewString(s))
		}
		return optional.Empty[base.Node]()
	})))
}
