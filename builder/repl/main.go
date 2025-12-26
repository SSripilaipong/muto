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
	program *replProgram.Wrapper
}

func New(lineReader replReader.LineReader, printer replProgram.Printer) Repl {
	globalMod := global.NewBaseModule()
	ports := portal.NewDefaultPortal()
	imported := builtin.NewBuiltinImportMapping(nil).Attach(ports)

	mod := module.BuildUserDefinedModuleFromBase(globalMod, stBase.NewModule(nil)).
		Attach(module.NewDependency(ports, imported))

	prog := replProgram.New(program.New(mod), printer, ports, imported)
	return Repl{
		Reader:   replReader.New(command.NewParser(prog), lineReader, printer),
		Executor: replExecutor.New(prog),
		program:  prog,
	}
}
