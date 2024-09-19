package main

import (
	"os"

	"github.com/SSripilaipong/muto/cmd/cli/cli"
)

func main() {
	if err := cli.New()(os.Args); err != nil {
		panic(err)
	}
}
