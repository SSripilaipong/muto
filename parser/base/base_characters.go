package base

import (
	"github.com/SSripilaipong/muto/common/fn"
)

var EqualSign = chRune('=')
var AtSign = chRune('@')
var BackSlash = chRune('\\')
var MinusSign = chRune('-')
var Dot = chRune('.')
var ThreeDots = fixedChars("...")
var OpenParenthesis = chRune('(')
var CloseParenthesis = chRune(')')
var DoubleQuote = char(IsDoubleQuote)
var NotDoubleQuote = char(fn.Not(IsDoubleQuote))
var Digit = char(IsDigit)
var Space = char(IsSpace)
var LineBreak = char(IsLineBreak)
