package base

import (
	psPred "github.com/SSripilaipong/muto/parser/predicate"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
)

var chEqualSign = char(psPred.IsEqualSign)
var chAtSign = char(psPred.IsAtSign)
var chOpenParenthesis = char(psPred.IsOpenParenthesis)
var chCloseParenthesis = char(psPred.IsCloseParenthesis)
var chDoubleQuote = char(psPred.IsDoubleQuote)
var chBackSlash = char(psPred.IsBackSlash)
var chDigit = char(psPred.IsFirstRuneDigit)
var chMinusSign = char(psPred.IsMinusSign)
var chDot = char(psPred.IsDot)

func tokenToRunes(x tk.Token) []rune {
	return []rune(x.Value())
}
