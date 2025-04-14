package base

import (
	"fmt"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Number = ps.Or(
	ps.Map(st.NewNumber, unsignedNumber),
	ps.Map(digitsWithMinusSign, ps.Sequence2(MinusSign, unsignedNumber)),
)

var NumberResultNode = ps.Map(stResult.ToNode, Number)

var NumberPattern = ps.Map(st.ToPattern, Number)

var digitsWithMinusSign = tuple.Fn2(func(ms Character, x string) st.Number {
	return st.NewNumber("-" + x)
})

var unsignedNumber = ps.First(
	ps.Map(floatingNumber, ps.Sequence3(digits, Dot, digits)),
	digits,
)

var digits = ps.Map(tokensToString, ps.GreedyRepeatAtLeastOnce(Digit))

var floatingNumber = tuple.Fn3(func(left string, _dot Character, right string) string {
	return fmt.Sprintf("%s.%s", left, right)
})
