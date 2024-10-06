package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
	tk "github.com/SSripilaipong/muto/parser/tokens"
)

var identifierStartingWithLowerCase = ps.Map(
	fn.Compose(tk.NewIdentifier, tuple.Fn2(joinTokenString)), ps.Sequence2(char(psPred.IsIdentifierFirstLetterLowerCase), identifierFollowingLetters),
)

var identifierStartingWithUpperCase = ps.Map(
	fn.Compose(tk.NewIdentifier, tuple.Fn2(joinTokenString)), ps.Sequence2(char(psPred.IsIdentifierFirstLetterUpperCase), identifierFollowingLetters),
)

var identifierFollowingLetters = ps.Map(tokensToString, ps.OptionalGreedyRepeat(char(psPred.IsIdentifierFollowingLetter)))
