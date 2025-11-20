package parsing

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
)

func char(x rune) func([]rune) tuple.Of2[rslt.Of[rune], []rune] {
	return ConsumeIf(func(s rune) bool { return s == x })
}
