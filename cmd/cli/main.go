package main

import (
	"os"

	"muto/cmd/cli/run"
)

func main() {
	cli := New(run.NewCommand())

	if err := cli(os.Args); err != nil {
		panic(err)
	}
}
