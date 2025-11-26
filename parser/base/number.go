package base

import (
	"fmt"

	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/syntaxtree"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Number = ps.First(
	ps.Map(syntaxtree.NewNumber, ps.ToParser(unsignedNumber)),
	ps.Map(digitsWithMinusSign, ps.Sequence2(
		ps.ToParser(MinusSign),
		ps.ToParser(unsignedNumber),
	)),
).Legacy

var NumberResultNode = ps.Map(stResult.ToNode, ps.ToParser(Number)).Legacy

var NumberPattern = ps.Map(st.ToPattern, ps.ToParser(Number)).Legacy

var digitsWithMinusSign = tuple.Fn2(func(ms Character, x string) syntaxtree.Number {
	return syntaxtree.NewNumber("-" + x)
})

var unsignedNumber = ps.First(
	ps.Map(floatingNumber, ps.Sequence3(
		ps.ToParser(digits),
		ps.ToParser(Dot),
		ps.ToParser(digits),
	)),
	ps.ToParser(digits),
).Legacy

var digits = ps.Map(tokensToString, ps.GreedyRepeatAtLeastOnce(ps.ToParser(Digit))).Legacy

var floatingNumber = tuple.Fn3(func(left string, _dot Character, right string) string {
	return fmt.Sprintf("%s.%s", left, right)
})
