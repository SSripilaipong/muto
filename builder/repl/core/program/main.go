package program

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/builtin"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/module"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/portal"
	"github.com/SSripilaipong/muto/program"
	"github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Wrapper struct {
	program       program.Program
	printer       Printer
	portal        *portal.Portal
	importMapping module.ImportMapping
}

func New(prog program.Program, printer Printer, portal *portal.Portal, importMapping module.ImportMapping) *Wrapper {
	return &Wrapper{
		program:       prog,
		printer:       printer,
		portal:        portal,
		importMapping: importMapping,
	}
}

func (w *Wrapper) AddRule(rule mutator.NamedUnit) optional.Of[int] {
	w.mainModule().AppendNormal(rule)
	return optional.Empty[int]()
}

func (w *Wrapper) ImportBuiltin(name string) optional.Of[int] {
	moduleNames := append(w.importMapping.Names(), name)
	w.importMapping = builtin.NewBuiltinImportMapping(moduleNames).Attach(w.portal)
	collection, found := w.importMapping.GetCollection(name).Return()
	if !found {
		panic("unexpected error")
	}
	w.mainModule().ExtendImportedCollection(name, collection)
	return optional.Empty[int]()
}

func (w *Wrapper) MutateNode(node base.Node) optional.Of[int] {
	result := w.program.MutateUntilTerminated(node)
	w.printer.Print(result.TopLevelString())
	return optional.Empty[int]()
}

func (w *Wrapper) BuildRule(rule syntaxtree.Rule) mutator.NamedUnit {
	return w.mainModule().BuildRule(rule)
}

func (w *Wrapper) BuildNode(object stResult.Object) optional.Of[base.Node] {
	return w.mainModule().BuildNode(object)
}

func (w *Wrapper) mainModule() module.Dynamic {
	return w.program.MainModule()
}
