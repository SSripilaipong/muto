package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/tuple"
)

func StringToCharTokens(s string) []Character {
	r := make([]Character, len(s))
	for i, x := range []rune(s) {
		r[i] = NewCharacter(x)
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

func char(g func(s rune) bool) Parser[Character] {
	return ps.ConsumeIf(func(x Character) bool {
		return g(x.Value())
	})
}

func chRune(x rune) Parser[Character] {
	return char(func(s rune) bool {
		return x == s
	})
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
