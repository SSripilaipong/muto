package repl

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/repl"
)

var command = ps.RsFirst(
	ps.RsMap(mergeQuitCommand, commandNamed("q")),
)

func mergeQuitCommand(string) repl.QuitCommand {
	return repl.NewQuitCommand()
}

func commandNamed(name string) base.RsParser[string] {
	return base.RsFixedChars(":" + name)
}
