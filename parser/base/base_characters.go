package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
)

var EqualSign = chRune('=')
var AtSign = chRune('@')
var BackSlash = chRune('\\')
var Slash = chRune('/')
var MinusSign = chRune('-')
var Dot = chRune('.')
var Comma = chRune(',')
var Colon = chRune(':')
var ThreeDots = FixedChars("...")
var OpenParenthesis = chRune('(')
var CloseParenthesis = chRune(')')
var OpenBrace = chRune('{')
var CloseBrace = chRune('}')
var OpenSquareBracket = chRune('[')
var CloseSquareBracket = chRune(']')
var DoubleQuote = char("double quote", IsDoubleQuote)
var NotDoubleQuote = char("not double quote", fn.Not(IsDoubleQuote))
var SingleQuote = char("single quote", IsSingleQuote)
var NotSingleQuote = char("not single quote", fn.Not(IsSingleQuote))
var Digit = char("digit", IsDigit)
var Space = char("space", IsSpace)
var LineBreak = char("line break", IsLineBreak)
var WhiteSpace = ps.First(Space, LineBreak)
var Alpha = char("alpha", fn.Or(fn.Or(IsASCIILetter, IsASCIIDigit), fn.Or(IsUnderscore, IsHyphen)))
