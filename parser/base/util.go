package base

import (
	"fmt"
	"strconv"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/slc"
)

func StringToCharTokens(s string) []Character {
	var lineNumber uint = 1
	var columnNumber uint = 1

	runes := []rune(s)
	r := make([]Character, len(runes))
	for i, x := range runes {
		r[i] = NewCharacter(x, lineNumber, columnNumber)

		columnNumber++
		if x == '\n' {
			lineNumber++
			columnNumber = 1
		}
	}
	return r
}

func FixedChars(s string) func([]Character) tuple.Of2[rslt.Of[string], []Character] {
	rs := []rune(s)
	n := len(rs)
	return func(xs []Character) tuple.Of2[rslt.Of[string], []Character] {
		if len(xs) < n {
			return tuple.New2(rslt.Error[string](fmt.Errorf("expected %s", s)), xs)
		}
		for i := range rs {
			if xs[i].Value() != rs[i] {
				return tuple.New2(rslt.Error[string](fmt.Errorf("expected %s", s)), xs)
			}
		}
		return tuple.New2(rslt.Value(s), xs[n:])
	}
}

func char(name string, g func(s rune) bool) func([]Character) tuple.Of2[rslt.Of[Character], []Character] {
	return func(xs []Character) tuple.Of2[rslt.Of[Character], []Character] {
		r := ps.ConsumeIf(func(x Character) bool { return g(x.Value()) })(xs)
		if r.X1().IsOk() {
			return r
		}
		return tuple.New2(rslt.Error[Character](fmt.Errorf("expected %s", name)), xs)
	}
}

func chRune(x rune) func([]Character) tuple.Of2[rslt.Of[Character], []Character] {
	return char(strconv.QuoteRune(x), func(s rune) bool { return x == s })
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
