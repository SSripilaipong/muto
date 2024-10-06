package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/parser/predicate"
	"github.com/SSripilaipong/muto/parser/tokenizer"
)

var symbol = ps.Map(fn.Compose(tokenizer.NewSymbol, tokensToString), ps.GreedyRepeatAtLeastOnce(char(predicate.IsSymbolLetter)))
