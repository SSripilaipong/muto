package base

import (
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
)

var identifierStartingWithNonUpperCase = ps.Map(
	tuple.Fn2(joinTokenString),
	ps.Sequence2(
		ps.ToParser(char("identifier first letter non upper case", IsIdentifierFirstLetterNonUpperCase)),
		identifierFollowingLetters,
	),
)

var identifierStartingWithUpperCaseAndUnderscore = ps.Map(
	tuple.Fn2(joinTokenString),
	ps.Sequence2(
		ps.ToParser(char("identifier first letter upper case or underscore", IsIdentifierFirstLetterUpperCaseAndUnderscore)),
		identifierFollowingLetters,
	),
)

var identifierStartingWithUpperCase = ps.Map(
	tuple.Fn2(joinTokenString),
	ps.Sequence2(
		ps.ToParser(char("identifier first letter upper case", IsIdentifierFirstLetterUpperCase)),
		identifierFollowingLetters,
	),
)

var identifierFollowingLetters = ps.Map(
	tokensToString,
	ps.OptionalGreedyRepeat(ps.ToParser(char("identifier following letter", IsIdentifierFollowingLetter))),
)
