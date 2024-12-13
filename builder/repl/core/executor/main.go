package executor

import (
	"github.com/SSripilaipong/muto/builder/repl/core/command"
	replProgram "github.com/SSripilaipong/muto/builder/repl/core/program"
	"github.com/SSripilaipong/muto/common/optional"
)

type Executor struct {
	program replProgram.Wrapper
}

func New(prog replProgram.Wrapper) Executor {
	return Executor{
		program: prog,
	}
}

func (e Executor) Execute(cmd command.Command) optional.Of[int] {
	switch {
	case command.IsMutateNodeCommand(cmd):
		return e.program.MutateNode(command.UnsafeCommandToMutateNodeCommand(cmd).InitialNode())
	case command.IsAddRuleCommand(cmd):
		return e.program.AddRule(command.UnsafeCommandToAddRuleCommand(cmd).Rule())
	}
	return optional.Empty[int]()
}
