package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/tuple"
)

func WithTrailingLineBreak[R any](p Parser[R]) Parser[R] {
	return ps.Map(
		tuple.Fn2(func(r R, _ []Character) R { return r }),
		ps.Sequence2(p, ps.OptionalGreedyRepeat(LineBreak)),
	)
}

func WithLeadingLineBreak[R any](p Parser[R]) Parser[R] {
	return ps.Map(
		tuple.Fn2(func(_ []Character, r R) R { return r }),
		ps.Sequence2(ps.OptionalGreedyRepeat(LineBreak), p.FunctionForm()),
	)
}

func IgnoreTrailingLineBreak[R any](p Parser[R]) Parser[R] {
	return ps.DrainTrailing(fn.Compose(IsLineBreak, CharacterToValue), p)
}

func IgnoreLeadingLineBreak[R any](p Parser[R]) Parser[R] {
	return ps.DrainLeading(fn.Compose(IsLineBreak, CharacterToValue), p)
}

func InParentheses[R any](x Parser[R]) Parser[R] {
	withoutParenthesis := func(x tuple.Of3[Character, R, Character]) R {
		return x.X2()
	}
	return ps.Map(withoutParenthesis, ps.Sequence3(OpenParenthesis, x, CloseParenthesis))
}

func InDoubleQuotes[R any](x Parser[R]) Parser[R] {
	withoutQuotes := func(x tuple.Of3[Character, R, Character]) R {
		return x.X2()
	}
	return ps.Map(withoutQuotes, ps.Sequence3(DoubleQuote, x, DoubleQuote))
}

func IgnoreSpaceBetween2[R1, R2 any](p1 Parser[R1], p2 Parser[R2]) Parser[tuple.Of2[R1, R2]] {
	merge := tuple.Fn3(func(r1 R1, _ []Character, r2 R2) tuple.Of2[R1, R2] {
		return tuple.New2(r1, r2)
	})
	return ps.Map(merge, ps.Sequence3(p1, ps.OptionalGreedyRepeat(Space), p2))
}

func SpaceSeparated2[R1, R2 any](p1 Parser[R1], p2 Parser[R2]) Parser[tuple.Of2[R1, R2]] {
	merge := tuple.Fn3(func(r1 R1, _ []Character, r2 R2) tuple.Of2[R1, R2] {
		return tuple.New2(r1, r2)
	})
	return ps.Map(merge, ps.Sequence3(p1, ps.GreedyRepeatAtLeastOnce(Space), p2))
}

func SpaceSeparated3[R1, R2, R3 any](p1 Parser[R1], p2 Parser[R2], p3 Parser[R3]) Parser[tuple.Of3[R1, R2, R3]] {
	merge := tuple.Fn5(func(r1 R1, _1 []Character, r2 R2, _2 []Character, r3 R3) tuple.Of3[R1, R2, R3] {
		return tuple.New3(r1, r2, r3)
	})
	return ps.Map(merge, ps.Sequence5(p1, ps.GreedyRepeatAtLeastOnce(Space), p2, ps.GreedyRepeatAtLeastOnce(Space), p3))
}

func GreedyRepeatAtLeastOnceSpaceSeparated[R any](p Parser[R]) Parser[[]R] {
	merge := tuple.Fn2(func(x R, xs []tuple.Of2[Character, R]) []R {
		return append([]R{x}, slc.Map(tuple.Of2ToX2[Character, R])(xs)...)
	})
	return ps.Map(merge, ps.Sequence2(p, ps.OptionalGreedyRepeat(ps.Sequence2(Space, p))))
}

func OptionalGreedyRepeatSpaceSeparated[R any](p Parser[R]) Parser[[]R] {
	return ps.First(
		GreedyRepeatAtLeastOnceSpaceSeparated(p),
		func(xs []Character) []tuple.Of2[[]R, []Character] { return SingleResult([]R{}, xs) },
	)
}
