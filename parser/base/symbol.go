package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
)

var symbol = ps.Map(tokensToString, ps.GreedyRepeatAtLeastOnce(ps.ToParser(char("symbol", IsSymbol))))
