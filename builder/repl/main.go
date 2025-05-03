package repl

import (
	"github.com/SSripilaipong/muto/builder/repl/core/command"
	replExecutor "github.com/SSripilaipong/muto/builder/repl/core/executor"
	replProgram "github.com/SSripilaipong/muto/builder/repl/core/program"
	replReader "github.com/SSripilaipong/muto/builder/repl/core/reader"
	"github.com/SSripilaipong/muto/core/mutation"
	"github.com/SSripilaipong/muto/core/mutation/builtin"
	"github.com/SSripilaipong/muto/program"
)

type Repl struct {
	replReader.Reader
	replExecutor.Executor
}

func New(lineReader replReader.LineReader, printer replProgram.Printer) Repl {
	pkg := mutation.NewPackageFromStatements(nil, builtin.NewBuiltinMutatorsForStdio())
	prog := replProgram.New(program.New(pkg), printer)

	return Repl{
		Reader:   replReader.New(command.NewParser(prog), lineReader, printer),
		Executor: replExecutor.New(prog),
	}
}
