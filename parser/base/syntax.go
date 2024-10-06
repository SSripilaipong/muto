package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
)

func WithTrailingLineBreak[R any](p Parser[R]) Parser[R] {
	return ps.Map(
		tuple.Fn2(func(r R, _ tk.Token) R { return r }),
		ps.Sequence2(p, ps.ConsumeIf(tk.IsLineBreak)),
	)
}

func WithLeadingLineBreak[R any](p Parser[R]) Parser[R] {
	return ps.Map(
		tuple.Fn2(func(_ tk.Token, r R) R { return r }),
		ps.Sequence2(ps.ConsumeIf(tk.IsLineBreak), p.FunctionForm()),
	)
}

func IgnoreTrailingLineBreak[R any](p Parser[R]) Parser[R] {
	return ps.DrainTrailing(tk.IsLineBreak, p)
}

func IgnoreLeadingLineBreak[R any](p Parser[R]) Parser[R] {
	return ps.DrainLeading(tk.IsLineBreak, p)
}

func InParentheses[R any](x Parser[R]) Parser[R] {
	withoutParenthesis := func(x tuple.Of3[tk.Token, R, tk.Token]) R {
		return x.X2()
	}
	return ps.Map(withoutParenthesis, ps.Sequence3(OpenParenthesis, x, CloseParenthesis))
}

func IgnoreSpaceBetween2[R1, R2 any](p1 Parser[R1], p2 Parser[R2]) Parser[tuple.Of2[R1, R2]] {
	merge := tuple.Fn3(func(r1 R1, _ []tk.Token, r2 R2) tuple.Of2[R1, R2] {
		return tuple.New2(r1, r2)
	})
	return ps.Map(merge, ps.Sequence3(p1, ps.OptionalGreedyRepeat(chSpace), p2))
}

func IgnoreSpaceBetween3[R1, R2, R3 any](p1 Parser[R1], p2 Parser[R2], p3 Parser[R3]) Parser[tuple.Of3[R1, R2, R3]] {
	merge := tuple.Fn5(func(r1 R1, _1 []tk.Token, r2 R2, _2 []tk.Token, r3 R3) tuple.Of3[R1, R2, R3] {
		return tuple.New3(r1, r2, r3)
	})
	return ps.Map(merge, ps.Sequence5(p1, ps.OptionalGreedyRepeat(chSpace), p2, ps.OptionalGreedyRepeat(chSpace), p3))
}

func SpaceSeparated2[R1, R2 any](p1 Parser[R1], p2 Parser[R2]) Parser[tuple.Of2[R1, R2]] {
	//merge := tuple.Fn3(func(r1 R1, _ []tk.Token, r2 R2) tuple.Of2[R1, R2] {
	//	return tuple.New2(r1, r2)
	//})
	//return ps.Map(merge, ps.Sequence3(p1, ps.GreedyRepeatAtLeastOnce(chSpace), p2))
	return IgnoreSpaceBetween2(p1, p2) // TODO replace
}

func SpaceSeparated3[R1, R2, R3 any](p1 Parser[R1], p2 Parser[R2], p3 Parser[R3]) Parser[tuple.Of3[R1, R2, R3]] {
	//merge := tuple.Fn5(func(r1 R1, _1 []tk.Token, r2 R2, _2 []tk.Token, r3 R3) tuple.Of3[R1, R2, R3] {
	//	return tuple.New3(r1, r2, r3)
	//})
	//return ps.Map(merge, ps.Sequence5(p1, ps.GreedyRepeatAtLeastOnce(chSpace), p2, ps.GreedyRepeatAtLeastOnce(chSpace), p3))
	return IgnoreSpaceBetween3(p1, p2, p3) // TODO replace
}
