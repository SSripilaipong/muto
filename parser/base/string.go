package base

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/tuple"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
	tk "github.com/SSripilaipong/muto/parser/tokens"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var String = ps.Map(
	fn.Compose(st.NewString, stringWithQuotes), InDoubleQuotes(innerString),
)

var StringResultNode = ps.Map(stringToResultNode, String)

var StringPatternParam = ps.Map(stringToPatternParam, String)

var innerString = ps.Map(slc.Flatten, ps.OptionalGreedyRepeat(ps.First(escapedStringChar, nonDoubleQuoteChar)))
var escapedStringChar = ps.Map(escapeStringCharToRunes, ps.Sequence2(chBackSlash, ps.ConsumeIf(fn.Const[tk.Token](true))))
var nonDoubleQuoteChar = ps.Map(tokenToRunes, char(fn.Not(psPred.IsDoubleQuote)))

func stringWithQuotes(x []rune) string {
	return fmt.Sprintf(`"%s"`, string(x))
}

var escapeStringCharToRunes = tuple.Fn2(func(_bs tk.Token, x tk.Token) []rune {
	return append([]rune{'\\'}, []rune(x.Value())...)
})

func stringToResultNode(x st.String) stResult.Node     { return x }
func stringToPatternParam(x st.String) stPattern.Param { return x }
