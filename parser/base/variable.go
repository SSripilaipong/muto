package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
)

var FixedVar = ps.Or(
	consumeId(psPred.IsVariableName),
	identifierStartingWithUpperCase,
)
