package repl

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/repl"
)

var command = ps.First(
	ps.Map(mergeQuitCommand, ps.ToParser(commandNamed("q"))),
	ps.ToParser(importCommand),
).Legacy

var importCommand = ps.Map(
	parseImportCommand,
	ps.ToParser(base.SpaceSeparated2(
		base.FixedChars(":import"),
		base.ImportPath,
	)),
).Legacy

var parseImportCommand = tuple.Fn2(func(_ string, path []string) repl.Statement {
	return repl.NewImportCommand(path)
})

func mergeQuitCommand(string) repl.Statement {
	return repl.NewQuitCommand()
}

func commandNamed(name string) func([]base.Character) tuple.Of2[rslt.Of[string], []base.Character] {
	return base.FixedChars(":" + name)
}
