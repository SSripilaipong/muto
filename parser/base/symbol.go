package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/parser/predicate"
	tk "github.com/SSripilaipong/muto/parser/tokens"
)

var symbol = ps.Map(fn.Compose(tk.NewSymbol, tokensToString), ps.GreedyRepeatAtLeastOnce(char(predicate.IsSymbolLetter)))
