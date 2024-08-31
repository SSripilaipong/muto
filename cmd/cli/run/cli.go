package run

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const explainFlag = "explain"

func NewCommand() *cli.Command {
	return &cli.Command{
		Name: "run",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: explainFlag + ",E",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				fmt.Println("error: filename required")
				return nil
			}
			var options []func(executeOptions) executeOptions
			if c.Bool(explainFlag) {
				options = append(options, withExplanation())
			}
			return ExecuteByFileName(c.Args().First(), options...)
		},
	}
}
