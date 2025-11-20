package repl

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/repl"
)

var command = ps.First(
	ps.Map(mergeQuitCommand, commandNamed("q")),
)

func mergeQuitCommand(string) repl.QuitCommand {
	return repl.NewQuitCommand()
}

func commandNamed(name string) func([]base.Character) tuple.Of2[rslt.Of[string], []base.Character] {
	return base.FixedChars(":" + name)
}
