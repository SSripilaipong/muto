package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
)

var EqualSign = char(psPred.IsEqualSign)
var AtSign = char(psPred.IsAtSign)
var OpenParenthesis = char(psPred.IsOpenParenthesis)
var CloseParenthesis = char(psPred.IsCloseParenthesis)
var ThreeDots = fixedChars("...")
var DoubleQuote = char(psPred.IsDoubleQuote)
var NotDoubleQuote = char(fn.Not(psPred.IsDoubleQuote))
var BackSlash = char(psPred.IsBackSlash)
var Digit = char(psPred.IsDigit)
var MinusSign = char(psPred.IsMinusSign)
var Dot = char(psPred.IsDot)
var Space = char(psPred.IsSpace)
var LineBreak = char(psPred.IsLineBreak)
