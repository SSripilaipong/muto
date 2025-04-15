package repl

import (
	replExecutor "github.com/SSripilaipong/muto/builder/repl/core/executor"
	replProgram "github.com/SSripilaipong/muto/builder/repl/core/program"
	replReader "github.com/SSripilaipong/muto/builder/repl/core/reader"
	"github.com/SSripilaipong/muto/core/mutation"
	"github.com/SSripilaipong/muto/core/mutation/normal/builtin"
	"github.com/SSripilaipong/muto/program"
)

type Repl struct {
	replReader.Reader
	replExecutor.Executor
}

func New(lineReader replReader.LineReader, printer replProgram.Printer) Repl {
	prog := replProgram.New(program.New(mutation.NewFromStatements(nil, builtin.NewBuiltinMutatorsForStdio())), printer)

	return Repl{
		Reader:   replReader.New(lineReader, printer),
		Executor: replExecutor.New(prog),
	}
}
