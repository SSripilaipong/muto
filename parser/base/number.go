package base

import (
	"fmt"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Number = ps.Or(
	ps.Map(st.NewNumber, unsignedNumber),
	ps.Map(digitsWithMinusSign, ps.Sequence2(MinusSign, unsignedNumber)),
)

var NumberResultNode = ps.Map(numberToResultNode, Number)

var NumberPatternParam = ps.Map(numberToPatternParam, Number)

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

func numberToPatternParam(x st.Number) stPattern.Param { return x }
func numberToResultNode(x st.Number) stResult.Node     { return x }
