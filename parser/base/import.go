package base

import (
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/slc"
)

var ImportPath = ps.Map(parseImportPath, ps.Sequence2(
	ps.ToParser(ImportPathToken),
	ps.OptionalGreedyRepeat(ps.ToParser(importSubPathToken)),
)).Legacy
var importSubPathToken = ps.Prefix(ps.ToParser(Slash), ps.ToParser(ImportPathToken)).Legacy

var ImportPathToken = ps.Map(
	CharactersToString,
	ps.GreedyRepeatAtLeastOnce(ps.ToParser(Alpha)),
).Legacy

var parseImportPath = tuple.Fn2(func(t string, ts []string) []string {
	return append(slc.Pure(t), ts...)
})
