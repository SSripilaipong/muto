package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
)

var Class = ps.Or(
	consumeId(psPred.IsClassName),
	consumeSymbol(psPred.IsSymbol),
	identifierStartingWithLowerCase,
	symbol,
)
