package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
)

var EqualSign = consumeValue(psPred.IsEqualSign)
var AtSign = consumeValue(psPred.IsAtSign)
var OpenParenthesis = consumeValue(psPred.IsOpenParenthesis)
var CloseParenthesis = consumeValue(psPred.IsCloseParenthesis)

var String = ps.ConsumeIf(tk.IsString)
var Number = ps.ConsumeIf(tk.IsNumber)
var Boolean = consumeId(psPred.IsBooleanValue)
var Class = consumeId(psPred.IsClassName)
var Symbol = consumeSymbol(psPred.IsNotEqualSign)
var ClassIncludingSymbols = ps.Or(Class, Symbol)

var Variable = consumeId(psPred.IsVariableName)
