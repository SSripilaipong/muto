package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
)

var String = ps.Or(
	ps.ConsumeIf(tk.IsString),
	ps.Map(mergeStringChars, ps.Sequence3(chDoubleQuote, innerString, chDoubleQuote)),
)
var innerString = ps.Map(innerStringToRunes, ps.OptionalGreedyRepeat(ps.First(escapedStringChar, nonDoubleQuoteChar)))
var escapedStringChar = ps.Map(escapeStringCharToRunes, ps.Sequence2(chBackSlash, ps.ConsumeIf(fn.Const[tk.Token](true))))
var nonDoubleQuoteChar = ps.Map(tokenToRunes, char(fn.Not(psPred.IsDoubleQuote)))

func innerStringToRunes(xs [][]rune) (y []rune) {
	for _, x := range xs {
		y = append(y, x...)
	}
	return
}

var mergeStringChars = tuple.Fn3(func(_ tk.Token, x []rune, _ tk.Token) tk.Token {
	return tk.NewString("\"" + string(x) + "\"")
})

var escapeStringCharToRunes = tuple.Fn2(func(_bs tk.Token, x tk.Token) []rune {
	return append([]rune{'\\'}, []rune(x.Value())...)
})
