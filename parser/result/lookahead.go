package result

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
)

func EOL[R any](p func([]psBase.Character) tuple.Of2[rslt.Of[R], []psBase.Character]) func([]psBase.Character) tuple.Of2[rslt.Of[R], []psBase.Character] {
	return ps.Lookahead(func(s []psBase.Character) bool {
		return len(s) == 0 || psBase.IsLineBreak(s[0].Value())
	}, ps.ToParser(p)).Legacy
}
