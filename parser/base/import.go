package base

import (
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/slc"
)

var RsImportPath = ps.RsMap(parseImportPath, ps.RsSequence2(RsImportPathToken, ps.RsOptionalGreedyRepeat(importSubPathToken)))
var importSubPathToken = ps.Prefix(Slash, RsImportPathToken)

var ImportPathToken = ps.DeRs(RsImportPathToken)
var RsImportPathToken = ps.RsMap(CharactersToString, ps.RsGreedyRepeatAtLeastOnce(Alpha))

var parseImportPath = tuple.Fn2(func(t string, ts []string) []string {
	return append(slc.Pure(t), ts...)
})
