package program

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	coreMutation "github.com/SSripilaipong/muto/core/mutation"
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

func (w Wrapper) AddRule(rule mutator.NamedObjectMutator) optional.Of[int] {
	w.mainPackage().AppendNormalRule(rule)
	return optional.Empty[int]()
}

func (w Wrapper) MutateNode(node base.Node) optional.Of[int] {
	result := w.program.MutateUntilTerminated(node)
	w.printer.Print(result.TopLevelString())
	return optional.Empty[int]()
}

func (w Wrapper) BuildRule(rule syntaxtree.Rule) mutator.NamedObjectMutator {
	return w.mainPackage().BuildRule(rule)
}

func (w Wrapper) BuildNode(object stResult.Object) optional.Of[base.Node] {
	return w.mainPackage().BuildNode(object)
}

func (w Wrapper) mainPackage() coreMutation.Package {
	return w.program.MainPackage()
}
