package main

import (
	"os"

	"phi-lang/cmd/cli/run"
)

func main() {
	cli := New(run.NewCommand())

	if err := cli(os.Args); err != nil {
		panic(err)
	}
}
