package file

import (
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree"
)

var Command = ps.RsFirst(importCommand)

var importCommand = ps.RsMap(parseImportCommand, psBase.RsSpaceSeparated2(psBase.RsFixedChars(":import"), psBase.RsImportPath))

var parseImportCommand = tuple.Fn2(func(_ string, path []string) syntaxtree.Import {
	return syntaxtree.NewImport(path)
})
