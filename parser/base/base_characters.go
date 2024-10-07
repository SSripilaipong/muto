package base

import (
	psPred "github.com/SSripilaipong/muto/parser/predicate"
	tk "github.com/SSripilaipong/muto/parser/tokens"
)

var EqualSign = char(psPred.IsEqualSign)
var AtSign = char(psPred.IsAtSign)
var OpenParenthesis = char(psPred.IsOpenParenthesis)
var CloseParenthesis = char(psPred.IsCloseParenthesis)
var ThreeDots = fixedChars("...")
var chDoubleQuote = char(psPred.IsDoubleQuote)
var chBackSlash = char(psPred.IsBackSlash)
var chDigit = char(psPred.IsFirstRuneDigit)
var chMinusSign = char(psPred.IsMinusSign)
var chDot = char(psPred.IsDot)
var chSpace = char(psPred.IsSpace)
var chLineBreak = char(psPred.IsLineBreak)

func tokenToRunes(x tk.Token) []rune {
	return []rune(x.Value())
}
