package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
)

var consumeValue = consumeTokenWithValueCondition(fn.Const[tk.Token](true))

var consumeId = consumeTokenWithValueCondition(tk.IsIdentifier)

var consumeSymbol = consumeTokenWithValueCondition(tk.IsSymbol)

func consumeTokenWithValueCondition(f func(x tk.Token) bool) func(g func(s string) bool) func([]tk.Token) []tuple.Of2[tk.Token, []tk.Token] {
	return func(g func(s string) bool) func([]tk.Token) []tuple.Of2[tk.Token, []tk.Token] {
		return ps.ConsumeIf(func(x tk.Token) bool {
			return f(x) && g(x.Value())
		})
	}
}
