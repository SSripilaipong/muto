package base

import (
	"fmt"

	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/syntaxtree"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Number = ps.Or(
	ps.Map(syntaxtree.NewNumber, unsignedNumber),
	ps.Map(digitsWithMinusSign, ps.Sequence2(MinusSign, unsignedNumber)),
)

var NumberResultNode = ps.Map(stResult.ToNode, Number)

var NumberPattern = ps.Map(st.ToPattern, Number)

var digitsWithMinusSign = tuple.Fn2(func(ms Character, x string) syntaxtree.Number {
	return syntaxtree.NewNumber("-" + x)
})

var unsignedNumber = ps.First(
	ps.Map(floatingNumber, ps.Sequence3(digits, Dot, digits)),
	digits,
)

var digits = ps.Map(tokensToString, ps.GreedyRepeatAtLeastOnce(Digit))

var floatingNumber = tuple.Fn3(func(left string, _dot Character, right string) string {
	return fmt.Sprintf("%s.%s", left, right)
})
