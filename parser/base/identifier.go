package base

import (
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
)

var identifierStartingWithNonUpperCase = ps.Map(
	tuple.Fn2(joinTokenString), ps.Sequence2(char(IsIdentifierFirstLetterNonUpperCase), identifierFollowingLetters),
)

var identifierStartingWithUpperCaseAndUnderscore = ps.Map(
	tuple.Fn2(joinTokenString), ps.Sequence2(char(IsIdentifierFirstLetterUpperCaseAndUnderscore), identifierFollowingLetters),
)

var identifierStartingWithUpperCase = ps.Map(
	tuple.Fn2(joinTokenString), ps.Sequence2(char(IsIdentifierFirstLetterUpperCase), identifierFollowingLetters),
)

var identifierFollowingLetters = ps.Map(tokensToString, ps.OptionalGreedyRepeat(char(IsIdentifierFollowingLetter)))
