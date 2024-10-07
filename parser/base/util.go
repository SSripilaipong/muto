package base

import (
	"strings"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/tuple"
	tk "github.com/SSripilaipong/muto/parser/tokens"
)

func StringToCharTokens(s string) []tk.Token {
	r := make([]tk.Token, len(s))
	for i, x := range []rune(s) {
		r[i] = tk.NewCharacter(x)
	}
	return r
}

var char = consumeTokenWithValueCondition(tk.IsCharacter)

func fixedChars(s string) Parser[string] {
	rs := []rune(s)
	n := len(rs)
	return func(xs []tk.Token) []tuple.Of2[string, []tk.Token] {
		if len(xs) < n {
			return nil
		}
		for i := range rs {
			x := xs[i]
			if !tk.IsCharacter(x) || []rune(x.Value())[0] != rs[i] {
				return nil
			}
		}
		return slc.Pure(tuple.New2(s, xs[n:]))
	}
}

func consumeTokenWithValueCondition(f func(x tk.Token) bool) func(g func(s string) bool) Parser[tk.Token] {
	return func(g func(s string) bool) Parser[tk.Token] {
		return ps.ConsumeIf(func(x tk.Token) bool {
			return f(x) && g(x.Value())
		})
	}
}

func tokensToString(xs []tk.Token) string {
	var ss []string
	for _, x := range xs {
		ss = append(ss, x.Value())
	}
	return strings.Join(ss, "")
}

var joinTokenString = func(x tk.Token, xs string) string {
	return x.Value() + xs
}
