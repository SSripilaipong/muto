package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
)

var FixedVar = ps.Or(
	consumeId(psPred.IsVariableName),
	ps.LookaheadCondition(fn.Not(ps.Matches(fixedChars("..."))), identifierStartingWithUpperCase),
)
