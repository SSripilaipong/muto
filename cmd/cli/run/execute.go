package run

import (
	"fmt"
	"os"

	programBuilder "github.com/SSripilaipong/muto/builder/program"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/program"
)

func ExecuteByFileName(fileName string, options ...func(executeOptions) executeOptions) error {
	execOpt := buildExecuteOptions(options)
	src, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf(`cannot read file "%s": %w`, fileName, err)
	}

	prog, err := buildWithOptions(src, execOpt)
	if err != nil {
		return fmt.Errorf(`cannot build from file "%s": %w`, fileName, err)
	}

	_ = prog.MutateUntilTerminated(prog.InitialObject())
	return nil
}

func buildWithOptions(src []byte, opt executeOptions) (program.Program, error) {
	prog, err := programBuilder.BuildProgramFromString(string(src)).Return()
	if err != nil {
		return program.Program{}, err
	}
	if opt.Explain {
		prog = prog.WithAfterMutationHook(func(node base.Node) { fmt.Println(node.TopLevelString()) })
	}
	return prog, nil
}

type executeOptions struct {
	Explain bool
}

func buildExecuteOptions(options []func(executeOptions) executeOptions) (opt executeOptions) {
	for _, option := range options {
		opt = option(opt)
	}
	return
}

func withExplanation() func(executeOptions) executeOptions {
	return func(opt executeOptions) executeOptions {
		opt.Explain = true
		return opt
	}
}
