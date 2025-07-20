package program

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	coreMutation "github.com/SSripilaipong/muto/core/module"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/program"
	"github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Wrapper struct {
	program program.Program
	printer Printer
}

func New(prog program.Program, printer Printer) Wrapper {
	return Wrapper{program: prog, printer: printer}
}

func (w Wrapper) AddRule(rule mutator.NamedUnit) optional.Of[int] {
	w.mainModule().AppendNormal(rule)
	return optional.Empty[int]()
}

func (w Wrapper) MutateNode(node base.Node) optional.Of[int] {
	result := w.program.MutateUntilTerminated(node)
	w.printer.Print(result.TopLevelString())
	return optional.Empty[int]()
}

func (w Wrapper) BuildRule(rule syntaxtree.Rule) mutator.NamedUnit {
	return w.mainModule().BuildRule(rule)
}

func (w Wrapper) BuildNode(object stResult.Object) optional.Of[base.Node] {
	return w.mainModule().BuildNode(object)
}

func (w Wrapper) mainModule() coreMutation.Dynamic {
	return w.program.MainModule()
}
