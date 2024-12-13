package cli

import (
	"github.com/urfave/cli/v2"

	"github.com/SSripilaipong/muto/cmd/cli/repl"
	"github.com/SSripilaipong/muto/cmd/cli/run"
)

func New() func(arguments []string) (err error) {
	return newWithCommands(run.NewCommand(), repl.NewCommand())
}

func newWithCommands(runCommand *cli.Command, replCommand *cli.Command) func(arguments []string) (err error) {
	app := &cli.App{
		Commands: []*cli.Command{
			runCommand,
			replCommand,
		},
	}
	return app.Run
}
