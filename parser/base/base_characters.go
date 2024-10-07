package base

import (
	"github.com/SSripilaipong/muto/common/fn"
)

var EqualSign = char(IsEqualSign)
var AtSign = char(IsAtSign)
var OpenParenthesis = char(IsOpenParenthesis)
var CloseParenthesis = char(IsCloseParenthesis)
var ThreeDots = fixedChars("...")
var DoubleQuote = char(IsDoubleQuote)
var NotDoubleQuote = char(fn.Not(IsDoubleQuote))
var BackSlash = char(IsBackSlash)
var Digit = char(IsDigit)
var MinusSign = char(IsMinusSign)
var Dot = char(IsDot)
var Space = char(IsSpace)
var LineBreak = char(IsLineBreak)
