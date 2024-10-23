package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/rslt"
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
		ps.Sequence2(ps.GreedyRepeatAtLeastOnce(LineBreak), p.FunctionForm()),
	)
}

func RsWithLeadingLineBreak[R any](p RsParser[R]) RsParser[R] {
	return ps.RsMap(
		tuple.Fn2(func(_ []Character, r R) R { return r }),
		ps.RsSequence2(ps.RsGreedyRepeatAtLeastOnce(RsLineBreak), p.FunctionForm()),
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

func InBracesWhiteSpacesAllowed[R any](x Parser[R]) Parser[R] {
	withoutBrace := func(x tuple.Of3[Character, R, Character]) R {
		return x.X2()
	}
	return ps.Map(withoutBrace, IgnoreWhiteSpaceBetween3(OpenBrace, x, CloseBrace))
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

func RsIgnoreSpaceBetween2[R1, R2 any](p1 RsParser[R1], p2 RsParser[R2]) RsParser[tuple.Of2[R1, R2]] {
	merge := tuple.Fn3(func(r1 R1, _ []Character, r2 R2) tuple.Of2[R1, R2] {
		return tuple.New2(r1, r2)
	})
	return ps.RsMap(merge, ps.RsSequence3(p1, ps.RsOptionalGreedyRepeat(RsSpace), p2))
}

func IgnoreSpaceBetween3[R1, R2, R3 any](p1 Parser[R1], p2 Parser[R2], p3 Parser[R3]) Parser[tuple.Of3[R1, R2, R3]] {
	merge := tuple.Fn2(func(r1 R1, r23 tuple.Of2[R2, R3]) tuple.Of3[R1, R2, R3] {
		return tuple.New3(r1, r23.X1(), r23.X2())
	})
	return ps.Map(merge, IgnoreSpaceBetween2(p1, IgnoreSpaceBetween2(p2, p3)))
}

func IgnoreWhiteSpaceBetween2[R1, R2 any](p1 Parser[R1], p2 Parser[R2]) Parser[tuple.Of2[R1, R2]] {
	merge := tuple.Fn3(func(r1 R1, _ []Character, r2 R2) tuple.Of2[R1, R2] {
		return tuple.New2(r1, r2)
	})
	return ps.Map(merge, ps.Sequence3(p1, ps.OptionalGreedyRepeat(WhiteSpace), p2))
}

func IgnoreWhiteSpaceBetween3[R1, R2, R3 any](p1 Parser[R1], p2 Parser[R2], p3 Parser[R3]) Parser[tuple.Of3[R1, R2, R3]] {
	merge := tuple.Fn2(func(r1 R1, r23 tuple.Of2[R2, R3]) tuple.Of3[R1, R2, R3] {
		return tuple.New3(r1, r23.X1(), r23.X2())
	})
	return ps.Map(merge, IgnoreWhiteSpaceBetween2(p1, IgnoreWhiteSpaceBetween2(p2, p3)))
}

func SpaceSeparated2[R1, R2 any](p1 Parser[R1], p2 Parser[R2]) Parser[tuple.Of2[R1, R2]] {
	merge := tuple.Fn3(func(r1 R1, _ []Character, r2 R2) tuple.Of2[R1, R2] {
		return tuple.New2(r1, r2)
	})
	return ps.Map(merge, ps.Sequence3(p1, ps.GreedyRepeatAtLeastOnce(Space), p2))
}

func RsSpaceSeparated2[R1, R2 any](p1 RsParser[R1], p2 RsParser[R2]) RsParser[tuple.Of2[R1, R2]] {
	merge := rslt.Fmap(tuple.Fn3(func(r1 R1, _ []Character, r2 R2) tuple.Of2[R1, R2] {
		return tuple.New2(r1, r2)
	}))
	return ps.Map(merge, ps.RsSequence3(p1, ps.RsGreedyRepeatAtLeastOnce(RsSpace), p2))
}

func SpaceSeparated3[R1, R2, R3 any](p1 Parser[R1], p2 Parser[R2], p3 Parser[R3]) Parser[tuple.Of3[R1, R2, R3]] {
	merge := tuple.Fn5(func(r1 R1, _1 []Character, r2 R2, _2 []Character, r3 R3) tuple.Of3[R1, R2, R3] {
		return tuple.New3(r1, r2, r3)
	})
	return ps.Map(merge, ps.Sequence5(p1, ps.GreedyRepeatAtLeastOnce(Space), p2, ps.GreedyRepeatAtLeastOnce(Space), p3))
}

func RsSpaceSeparated3[R1, R2, R3 any](p1 RsParser[R1], p2 RsParser[R2], p3 RsParser[R3]) RsParser[tuple.Of3[R1, R2, R3]] {
	merge := rslt.Fmap(tuple.Fn5(func(r1 R1, _1 []Character, r2 R2, _2 []Character, r3 R3) tuple.Of3[R1, R2, R3] {
		return tuple.New3(r1, r2, r3)
	}))
	return ps.Map(merge, ps.RsSequence5(p1, ps.RsGreedyRepeatAtLeastOnce(RsSpace), p2, ps.RsGreedyRepeatAtLeastOnce(RsSpace), p3))
}

func WhiteSpaceSeparated2[R1, R2 any](p1 Parser[R1], p2 Parser[R2]) Parser[tuple.Of2[R1, R2]] {
	merge := tuple.Fn3(func(r1 R1, _ []Character, r2 R2) tuple.Of2[R1, R2] {
		return tuple.New2(r1, r2)
	})
	return ps.Map(merge, ps.Sequence3(p1, ps.GreedyRepeatAtLeastOnce(WhiteSpace), p2))
}

func WhiteSpaceSeparated3[R1, R2, R3 any](p1 Parser[R1], p2 Parser[R2], p3 Parser[R3]) Parser[tuple.Of3[R1, R2, R3]] {
	merge := tuple.Fn5(func(r1 R1, _1 []Character, r2 R2, _2 []Character, r3 R3) tuple.Of3[R1, R2, R3] {
		return tuple.New3(r1, r2, r3)
	})
	return ps.Map(merge, ps.Sequence5(p1, ps.GreedyRepeatAtLeastOnce(WhiteSpace), p2, ps.GreedyRepeatAtLeastOnce(WhiteSpace), p3))
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

func GreedyRepeatAtLeastOnceIgnoreWhiteSpaceBetween[R any](p Parser[R]) Parser[[]R] {
	merge := tuple.Fn2(func(x R, xs []R) []R {
		return append([]R{x}, xs...)
	})
	return ps.Map(merge, ps.Sequence2(p, ps.OptionalGreedyRepeat(OptionalLeadingWhiteSpace(p))))
}

func OptionalGreedyRepeatIgnoreWhiteSpaceBetween[R any](p Parser[R]) Parser[[]R] {
	return ps.First(
		GreedyRepeatAtLeastOnceIgnoreWhiteSpaceBetween(p),
		func(xs []Character) []tuple.Of2[[]R, []Character] { return SingleResult([]R{}, xs) },
	)
}

func EndingWithCommaSpaceAllowed[R any](p Parser[R]) Parser[R] {
	return ps.Map(tuple.Of2ToX1, IgnoreSpaceBetween2(p, Comma))
}

func OptionalLeadingWhiteSpace[R any](p Parser[R]) Parser[R] {
	return ps.Map(tuple.Of2ToX2, ps.Sequence2(ps.OptionalGreedyRepeat(WhiteSpace), p))
}
