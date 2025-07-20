package repl

import (
	"github.com/SSripilaipong/muto/builder/repl/core/command"
	replExecutor "github.com/SSripilaipong/muto/builder/repl/core/executor"
	replProgram "github.com/SSripilaipong/muto/builder/repl/core/program"
	replReader "github.com/SSripilaipong/muto/builder/repl/core/reader"
	"github.com/SSripilaipong/muto/builtin"
	"github.com/SSripilaipong/muto/builtin/global"
	"github.com/SSripilaipong/muto/builtin/portal"
	"github.com/SSripilaipong/muto/core/module"
	"github.com/SSripilaipong/muto/program"
	stBase "github.com/SSripilaipong/muto/syntaxtree"
)

type Repl struct {
	replReader.Reader
	replExecutor.Executor
}

func New(lineReader replReader.LineReader, printer replProgram.Printer) Repl {
	globalMod := global.NewModule()
	ports := portal.NewDefaultPortal()
	imported := builtin.NewBuiltinImportMapping(nil).Attach(globalMod, ports)

	mod := module.BuildUserDefinedModule(stBase.NewModule(nil)).
		Attach(module.NewDependency(globalMod, ports, imported))

	prog := replProgram.New(program.New(mod), printer)
	return Repl{
		Reader:   replReader.New(command.NewParser(prog), lineReader, printer),
		Executor: replExecutor.New(prog),
	}
}
