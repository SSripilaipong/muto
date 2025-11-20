package base

import (
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/slc"
)

var ImportPath = ps.Map(parseImportPath, ps.Sequence2(
	ImportPathToken,
	ps.OptionalGreedyRepeat(importSubPathToken),
))
var importSubPathToken = ps.Prefix(Slash, ImportPathToken)

var ImportPathToken = ps.Map(CharactersToString, ps.GreedyRepeatAtLeastOnce(Alpha))

var parseImportPath = tuple.Fn2(func(t string, ts []string) []string {
	return append(slc.Pure(t), ts...)
})
