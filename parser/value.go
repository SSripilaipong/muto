package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/parser/tokenizer"
)

var string_ = ps.ConsumeIf(tokenizer.IsString)
var number = ps.ConsumeIf(tokenizer.IsNumber)
var equalSign = ps.ConsumeIf(isEqualSign)
var atSign = ps.ConsumeIf(isAtSign)
var openParenthesis = ps.ConsumeIf(isOpenParenthesis)
var closeParenthesis = ps.ConsumeIf(isCloseParenthesis)
var symbol = ps.ConsumeIf(tokenizer.IsSymbol)
var identifier = ps.ConsumeIf(tokenizer.IsIdentifier)
var nonKeywordNonCapitalIdentifier = ps.ConsumeIf(func(x tokenizer.Token) bool {
	return tokenizer.IsIdentifier(x) && !isFirstLetterCapital(x.Value()) && !isKeyword(x.Value())
})
var symbolName = ps.ConsumeIf(func(x tokenizer.Token) bool {
	return tokenizer.IsSymbol(x) && x.Value() != "="
})
var variable = ps.ConsumeIf(func(x tokenizer.Token) bool {
	name := x.Value()
	return tokenizer.IsIdentifier(x) && isFirstLetterCapital(name) && noVarSuffix(name)
})
var boolean = ps.ConsumeIf(func(x tokenizer.Token) bool {
	return tokenizer.IsIdentifier(x) && isBooleanValue(x.Value())
})
