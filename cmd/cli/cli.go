package main

import (
	"github.com/urfave/cli/v2"
)

func New(runCommand *cli.Command) func(arguments []string) (err error) {
	app := &cli.App{
		Commands: []*cli.Command{
			runCommand,
		},
	}
	return app.Run
}
