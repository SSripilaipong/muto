package run

import (
	"fmt"
	"os"

	"muto/builder"
)

func ExecuteByFileName(fileName string) error {
	src, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf(`cannot read file "%s": %w`, fileName, err)
	}

	program, err := builder.BuildFromString(string(src)).Return()
	if err != nil {
		return fmt.Errorf(`cannot build from file "%s": %w`, fileName, err)
	}

	result := program.MutateUntilTerminated(program.InitialObject())
	fmt.Println(result)
	return nil
}
