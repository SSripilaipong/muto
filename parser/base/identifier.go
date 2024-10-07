package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
)

var identifierStartingWithLowerCase = ps.Map(
	tuple.Fn2(joinTokenString), ps.Sequence2(char(psPred.IsIdentifierFirstLetterLowerCase), identifierFollowingLetters),
)

var identifierStartingWithUpperCase = ps.Map(
	tuple.Fn2(joinTokenString), ps.Sequence2(char(psPred.IsIdentifierFirstLetterUpperCase), identifierFollowingLetters),
)

var identifierFollowingLetters = ps.Map(tokensToString, ps.OptionalGreedyRepeat(char(psPred.IsIdentifierFollowingLetter)))
