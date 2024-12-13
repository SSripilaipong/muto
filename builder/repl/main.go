package repl

import (
	replExecutor "github.com/SSripilaipong/muto/builder/repl/core/executor"
	replProgram "github.com/SSripilaipong/muto/builder/repl/core/program"
	replReader "github.com/SSripilaipong/muto/builder/repl/core/reader"
	"github.com/SSripilaipong/muto/core/mutation"
	"github.com/SSripilaipong/muto/program"
)

type Repl struct {
	replReader.Reader
	replExecutor.Executor
}

func New() Repl {
	prog := replProgram.New(program.New(mutation.NewFromStatements(nil)))

	return Repl{
		Reader:   replReader.New(),
		Executor: replExecutor.New(prog),
	}
}
