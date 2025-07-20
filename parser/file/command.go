package file

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

var Command = ps.RsFirst(importCommand)

var importCommand = ps.RsMap(parseImportCommand, psBase.RsSpaceSeparated2(psBase.RsFixedChars(":import"), importPath))

var importPath = ps.RsMap(parseImportPath, ps.RsSequence2(importPathToken, ps.RsOptionalGreedyRepeat(importSubPathToken)))

var importSubPathToken = ps.Prefix(psBase.Dot, importPathToken)

var importPathToken = ps.RsMap(psBase.CharactersToString, ps.RsGreedyRepeatAtLeastOnce(psBase.Alpha))

var parseImportCommand = tuple.Fn2(func(_ string, path []string) base.Import {
	return base.NewImport(path)
})

var parseImportPath = tuple.Fn2(func(t string, ts []string) []string {
	return append(slc.Pure(t), ts...)
})
