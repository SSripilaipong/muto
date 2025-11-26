package file

import (
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree"
)

var Command = ps.First(ps.ToParser(importCommand))

var importCommand = ps.Map(
	parseImportCommand,
	ps.ToParser(psBase.SpaceSeparated2(
		psBase.FixedChars(":import"),
		psBase.ImportPath,
	)),
).Legacy

var parseImportCommand = tuple.Fn2(func(_ string, path []string) syntaxtree.Import {
	return syntaxtree.NewImport(path)
})
