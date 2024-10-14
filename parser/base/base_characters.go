package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
)

var EqualSign = chRune('=')
var AtSign = chRune('@')
var BackSlash = chRune('\\')
var MinusSign = chRune('-')
var Dot = chRune('.')
var Comma = chRune(',')
var Colon = chRune(':')
var ThreeDots = fixedChars("...")
var OpenParenthesis = chRune('(')
var CloseParenthesis = chRune(')')
var OpenBrace = chRune('{')
var CloseBrace = chRune('}')
var DoubleQuote = char(IsDoubleQuote)
var NotDoubleQuote = char(fn.Not(IsDoubleQuote))
var Digit = char(IsDigit)
var Space = char(IsSpace)
var LineBreak = char(IsLineBreak)
var WhiteSpace = ps.Or(Space, LineBreak)
