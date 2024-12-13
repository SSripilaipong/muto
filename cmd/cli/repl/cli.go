package repl

import (
	"os"

	"github.com/urfave/cli/v2"

	replBuilder "github.com/SSripilaipong/muto/builder/repl"
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:  "repl",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			os.Exit(loop(replBuilder.New()))
			return nil
		},
	}
}
