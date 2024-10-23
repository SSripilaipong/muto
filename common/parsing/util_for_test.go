package parsing

import (
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/tuple"
)

func char(x rune) func([]rune) []tuple.Of2[rslt.Of[rune], []rune] {
	return RsConsumeIf(func(s rune) bool { return s == x })
}
