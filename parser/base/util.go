package base

import (
	"fmt"
	"strconv"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/tuple"
)

func StringToCharTokens(s string) []Character {
	var lineNumber uint = 1
	var columnNumber uint = 1

	r := make([]Character, len(s))
	for i, x := range []rune(s) {
		r[i] = NewCharacter(x, lineNumber, columnNumber)

		columnNumber++
		if x == '\n' {
			lineNumber++
			columnNumber = 1
		}
	}
	return r
}

func Prefix[P, T any](pf Parser[P], x Parser[T]) Parser[T] {
	var merge = tuple.Fn2(func(_ P, x T) T {
		return x
	})
	return ps.Map(merge, ps.Sequence2(pf, x))
}

func fixedChars(s string) Parser[string] {
	rs := []rune(s)
	n := len(rs)
	return func(xs []Character) []tuple.Of2[string, []Character] {
		if len(xs) < n {
			return nil
		}
		for i := range rs {
			x := xs[i]
			if x.Value() != rs[i] {
				return nil
			}
		}
		return slc.Pure(tuple.New2(s, xs[n:]))
	}
}

func RsFixedChars(x string) RsParser[string] {
	return ps.RsFirst(
		ps.Map(rslt.Value, fixedChars(x)),
		ps.Result[Character](rslt.Error[string](fmt.Errorf("expected %s", x))),
	)
}

func char(g func(s rune) bool) Parser[Character] {
	return ps.ConsumeIf(func(x Character) bool {
		return g(x.Value())
	})
}

func rsChar(name string, g func(s rune) bool) RsParser[Character] {
	return ps.RsFirst(
		ps.Map(rslt.Value, char(g)),
		ps.Result[Character](rslt.Error[Character](fmt.Errorf("expected %s", name))),
	)
}

func chRune(x rune) Parser[Character] {
	return char(func(s rune) bool {
		return x == s
	})
}

func rsChRune(x rune) RsParser[Character] {
	return ps.RsFirst(
		ps.Map(rslt.Value, chRune(x)),
		ps.Result[Character](rslt.Error[Character](fmt.Errorf("expected %s", strconv.QuoteRune(x)))),
	)
}

func tokensToString(xs []Character) string {
	return string(slc.Map(CharacterToValue)(xs))
}

func tokenToRunes(x Character) []rune {
	return slc.Pure(x.Value())
}

var joinTokenString = func(x Character, xs string) string {
	return string(x.Value()) + xs
}
