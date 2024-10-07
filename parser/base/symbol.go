package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/parser/predicate"
)

var symbol = ps.Map(tokensToString, ps.GreedyRepeatAtLeastOnce(char(predicate.IsSymbolLetter)))
