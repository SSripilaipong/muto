package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/tuple"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
	tk "github.com/SSripilaipong/muto/parser/tokens"
)

func WithTrailingLineBreak[R any](p Parser[R]) Parser[R] {
	return ps.Map(
		tuple.Fn2(func(r R, _ []tk.Token) R { return r }),
		ps.Sequence2(p, ps.OptionalGreedyRepeat(chLineBreak)),
	)
}

func WithLeadingLineBreak[R any](p Parser[R]) Parser[R] {
	return ps.Map(
		tuple.Fn2(func(_ []tk.Token, r R) R { return r }),
		ps.Sequence2(ps.OptionalGreedyRepeat(chLineBreak), p.FunctionForm()),
	)
}

func IgnoreTrailingLineBreak[R any](p Parser[R]) Parser[R] {
	return ps.DrainTrailing(fn.Compose(psPred.IsLineBreak, tk.TokenToValue), p)
}

func IgnoreLeadingLineBreak[R any](p Parser[R]) Parser[R] {
	return ps.DrainLeading(fn.Compose(psPred.IsLineBreak, tk.TokenToValue), p)
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

func SpaceSeparated2[R1, R2 any](p1 Parser[R1], p2 Parser[R2]) Parser[tuple.Of2[R1, R2]] {
	merge := tuple.Fn3(func(r1 R1, _ []tk.Token, r2 R2) tuple.Of2[R1, R2] {
		return tuple.New2(r1, r2)
	})
	return ps.Map(merge, ps.Sequence3(p1, ps.GreedyRepeatAtLeastOnce(chSpace), p2))
}

func SpaceSeparated3[R1, R2, R3 any](p1 Parser[R1], p2 Parser[R2], p3 Parser[R3]) Parser[tuple.Of3[R1, R2, R3]] {
	merge := tuple.Fn5(func(r1 R1, _1 []tk.Token, r2 R2, _2 []tk.Token, r3 R3) tuple.Of3[R1, R2, R3] {
		return tuple.New3(r1, r2, r3)
	})
	return ps.Map(merge, ps.Sequence5(p1, ps.GreedyRepeatAtLeastOnce(chSpace), p2, ps.GreedyRepeatAtLeastOnce(chSpace), p3))
}

func GreedyRepeatAtLeastOnceSpaceSeparated[R any](p Parser[R]) Parser[[]R] {
	merge := tuple.Fn2(func(x R, xs []tuple.Of2[tk.Token, R]) []R {
		return append([]R{x}, slc.Map(tuple.Of2ToX2[tk.Token, R])(xs)...)
	})
	return ps.Map(merge, ps.Sequence2(p, ps.OptionalGreedyRepeat(ps.Sequence2(chSpace, p))))
}

func OptionalGreedyRepeatSpaceSeparated[R any](p Parser[R]) Parser[[]R] {
	return ps.First(
		GreedyRepeatAtLeastOnceSpaceSeparated(p),
		func(xs []tk.Token) []tuple.Of2[[]R, []tk.Token] { return SingleResult([]R{}, xs) },
	)
}
