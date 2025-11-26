package base

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/syntaxtree"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Rune = ps.Map(
	fn.Compose(syntaxtree.NewRune, stringWithSingleQuotes),
	ps.ToParser(InSingleQuotes(innerRune)),
).Legacy

var RuneResultNode = ps.Map(stResult.ToNode, ps.ToParser(Rune)).Legacy

var RunePattern = ps.Map(st.ToPattern, ps.ToParser(Rune)).Legacy

var innerRune = ps.First(
	ps.ToParser(escapedRune),
	ps.ToParser(nonEscapedRune),
).Legacy
var escapedRune = ps.Map(escapeStringToRunes,
	ps.Sequence2(
		ps.ToParser(BackSlash),
		ps.ToParser(ps.ConsumeIf(fn.Const[Character](true))),
	),
).Legacy
var nonEscapedRune = ps.Map(tokenToRunes, ps.ToParser(NotSingleQuote)).Legacy

func stringWithSingleQuotes(x []rune) string {
	return fmt.Sprintf(`'%s'`, string(x))
}
