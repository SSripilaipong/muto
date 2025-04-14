package base

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/tuple"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var String = ps.Map(
	fn.Compose(st.NewString, stringWithQuotes), InDoubleQuotes(innerString),
)

var StringResultNode = ps.Map(stResult.ToNode, String)

var StringPattern = ps.Map(st.ToPattern, String)

var innerString = ps.Map(slc.Flatten, ps.OptionalGreedyRepeat(ps.First(escapedChar, nonEscapedChar)))
var escapedChar = ps.Map(escapeStringCharToRunes, ps.Sequence2(BackSlash, ps.ConsumeIf(fn.Const[Character](true))))
var nonEscapedChar = ps.Map(tokenToRunes, NotDoubleQuote)

func stringWithQuotes(x []rune) string {
	return fmt.Sprintf(`"%s"`, string(x))
}

var escapeStringCharToRunes = tuple.Fn2(func(_bs Character, x Character) []rune {
	return []rune{'\\', x.Value()}
})
