package parser

import (
	ps "phi-lang/common/parsing"
	"phi-lang/tokenizer"
)

var string_ = ps.ConsumeIf(tokenizer.IsString)
var number = ps.ConsumeIf(tokenizer.IsNumber)
var equalSign = ps.ConsumeIf(isEqualSign)
var symbol = ps.ConsumeIf(tokenizer.IsSymbol)
var identifier = ps.ConsumeIf(tokenizer.IsIdentifier)
var nonCapitalIdentifier = ps.ConsumeIf(func(x tokenizer.Token) bool {
	return tokenizer.IsIdentifier(x) && !isFirstLetterCapital(x.Value())
})
var variable = ps.ConsumeIf(func(x tokenizer.Token) bool {
	return tokenizer.IsIdentifier(x) && isFirstLetterCapital(x.Value())
})
