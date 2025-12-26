package executor

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/builder/repl/core/command"
	replProgram "github.com/SSripilaipong/muto/builder/repl/core/program"
)

type Executor struct {
	program *replProgram.Wrapper
}

func New(prog *replProgram.Wrapper) Executor {
	return Executor{
		program: prog,
	}
}

func (e Executor) Execute(cmd command.Command) optional.Of[int] {
	switch {
	case command.IsQuitCommand(cmd):
		return optional.Value(0)
	case command.IsMutateNodeCommand(cmd):
		return e.program.MutateNode(command.UnsafeCommandToMutateNodeCommand(cmd).InitialNode())
	case command.IsAddRuleCommand(cmd):
		return e.program.AddRule(command.UnsafeCommandToAddRuleCommand(cmd).Rule())
	case command.IsImportCommand(cmd):
		return e.program.ImportBuiltin(command.UnsafeCommandToImportCommand(cmd).Name())
	}
	return optional.Empty[int]()
}
