package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
)

var EqualSign = ps.Or(
	consumeValue(psPred.IsEqualSign),
	chEqualSign,
)
var AtSign = ps.Or(
	consumeValue(psPred.IsAtSign),
	chAtSign,
)
var OpenParenthesis = ps.Or(
	consumeValue(psPred.IsOpenParenthesis),
	chOpenParenthesis,
)
var CloseParenthesis = ps.Or(
	consumeValue(psPred.IsCloseParenthesis),
	chCloseParenthesis,
)
