package base

import (
	"fmt"

	"github.com/SSripilaipong/go-common/tuple"

	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/syntaxtree"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var String = ps.Map(
	fn.Compose(syntaxtree.NewString, stringWithQuotes), ps.ToParser(InDoubleQuotes(innerString)),
).Legacy

var StringResultNode = ps.Map(stResult.ToNode, ps.ToParser(String)).Legacy

var StringPattern = ps.Map(st.ToPattern, ps.ToParser(String)).Legacy

var innerString = ps.Map(
	slc.Flatten,
	ps.OptionalGreedyRepeat(ps.First(
		ps.ToParser(escapedStringCharacter),
		ps.ToParser(nonEscapedStringCharacter),
	)),
).Legacy
var escapedStringCharacter = ps.Map(escapeStringToRunes,
	ps.Sequence2(
		ps.ToParser(BackSlash),
		ps.ToParser(ps.ConsumeIf(fn.Const[Character](true))),
	),
).Legacy
var nonEscapedStringCharacter = ps.Map(tokenToRunes, ps.ToParser(NotDoubleQuote)).Legacy

func stringWithQuotes(x []rune) string {
	return fmt.Sprintf(`"%s"`, string(x))
}

var escapeStringToRunes = tuple.Fn2(func(_bs Character, x Character) []rune {
	return []rune{'\\', x.Value()}
})
