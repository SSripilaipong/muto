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
	fn.Compose(syntaxtree.NewRune, stringWithSingleQuotes), InSingleQuotes(innerRune),
)

var RuneResultNode = ps.Map(stResult.ToNode, Rune)

var RunePattern = ps.Map(st.ToPattern, Rune)

var innerRune = ps.First(escapedRune, nonEscapedRune)
var escapedRune = ps.Map(escapeStringToRunes, ps.Sequence2(BackSlash, ps.ConsumeIf(fn.Const[Character](true))))
var nonEscapedRune = ps.Map(tokenToRunes, NotSingleQuote)

func stringWithSingleQuotes(x []rune) string {
	return fmt.Sprintf(`'%s'`, string(x))
}
