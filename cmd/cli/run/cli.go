package run

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name: "run",
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				fmt.Println("error: filename required")
				return nil
			}
			return ExecuteByFileName(c.Args().First())
		},
	}
}
