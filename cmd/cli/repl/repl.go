package repl

import (
	"github.com/SSripilaipong/go-common/optional"
)

type repl[Cmd any] interface {
	Read() optional.Of[Cmd]
	Execute(Cmd) optional.Of[int]
}

func loop[Cmd any](p repl[Cmd]) int {
	for {
		cmd := p.Read()
		if cmd.IsEmpty() {
			continue
		}

		exit := p.Execute(cmd.Value())

		if exit.IsNotEmpty() {
			return exit.Value()
		}
	}
}
