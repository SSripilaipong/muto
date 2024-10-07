package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
)

var symbol = ps.Map(tokensToString, ps.GreedyRepeatAtLeastOnce(char(psPred.IsSymbol)))
