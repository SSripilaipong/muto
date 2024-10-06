package base

import (
	"fmt"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	tk "github.com/SSripilaipong/muto/parser/tokens"
)

var Number = ps.Or(
	ps.ConsumeIf(tk.IsNumber),
	ps.Map(tk.NewNumber, unsignedNumber),
	ps.Map(digitsWithMinusSign, ps.Sequence2(chMinusSign, unsignedNumber)),
)

var digitsWithMinusSign = tuple.Fn2(func(ms tk.Token, x string) tk.Token {
	return tk.NewNumber("-" + x)
})

var unsignedNumber = ps.First(
	ps.Map(floatingNumber, ps.Sequence3(digits, chDot, digits)),
	digits,
)

var floatingNumber = tuple.Fn3(func(left string, _dot tk.Token, right string) string {
	return fmt.Sprintf("%s.%s", left, right)
})

var digits = ps.Map(tokensToString, ps.GreedyRepeatAtLeastOnce(chDigit))
