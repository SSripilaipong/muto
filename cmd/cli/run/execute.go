package run

import (
	"fmt"
	"os"

	"github.com/SSripilaipong/muto/builder"
	"github.com/SSripilaipong/muto/core/base"
)

func ExecuteByFileName(fileName string, options ...func(executeOptions) executeOptions) error {
	execOpt := buildExecuteOptions(options)
	src, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf(`cannot read file "%s": %w`, fileName, err)
	}

	program, err := buildWithOptions(src, execOpt)
	if err != nil {
		return fmt.Errorf(`cannot build from file "%s": %w`, fileName, err)
	}

	result := program.MutateUntilTerminated(program.InitialObject())
	if !execOpt.Explain {
		fmt.Println(result.TopLevelString())
	}
	return nil
}

func buildWithOptions(src []byte, opt executeOptions) (builder.Program, error) {
	program, err := builder.BuildFromString(string(src)).Return()
	if err != nil {
		return builder.Program{}, err
	}
	if opt.Explain {
		program = program.WithAfterMutationHook(func(node base.Node) { fmt.Println(node.TopLevelString()) })
	}
	return program, nil
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
