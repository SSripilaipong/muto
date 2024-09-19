package cli

import (
	"github.com/urfave/cli/v2"

	"github.com/SSripilaipong/muto/cmd/cli/run"
)

func New() func(arguments []string) (err error) {
	return newWithCommands(run.NewCommand())
}

func newWithCommands(runCommand *cli.Command) func(arguments []string) (err error) {
	app := &cli.App{
		Commands: []*cli.Command{
			runCommand,
		},
	}
	return app.Run
}
