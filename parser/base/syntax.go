package base

import (
	"errors"
	"math"
	"slices"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/slc"
)

func WithTrailingLineBreak[R any](p func([]Character) tuple.Of2[rslt.Of[R], []Character]) func([]Character) tuple.Of2[rslt.Of[R], []Character] {
	return ps.Map(
		tuple.Fn2(func(r R, _ []Character) R { return r }),
		ps.Sequence2(p, ps.OptionalGreedyRepeat(LineBreak)),
	)
}

func WithLeadingLineBreak[R any](p func([]Character) tuple.Of2[rslt.Of[R], []Character]) func([]Character) tuple.Of2[rslt.Of[R], []Character] {
	return ps.Map(
		tuple.Fn2(func(_ []Character, r R) R { return r }),
		ps.Sequence2(ps.GreedyRepeatAtLeastOnce(LineBreak), p),
	)
}

func IgnoreTrailingLineBreak[R any](p func([]Character) tuple.Of2[rslt.Of[R], []Character]) func([]Character) tuple.Of2[rslt.Of[R], []Character] {
	return ps.Map(
		tuple.Fn2(func(r R, _ []Character) R { return r }),
		ps.Sequence2(p, ps.OptionalGreedyRepeat(LineBreak)),
	)
}

func IgnoreLeadingLineBreak[R any](p func([]Character) tuple.Of2[rslt.Of[R], []Character]) func([]Character) tuple.Of2[rslt.Of[R], []Character] {
	return ps.Map(
		tuple.Fn2(func(_ []Character, r R) R { return r }),
		ps.Sequence2(ps.OptionalGreedyRepeat(LineBreak), p),
	)
}

func InParentheses[R any](x func([]Character) tuple.Of2[rslt.Of[R], []Character]) func([]Character) tuple.Of2[rslt.Of[R], []Character] {
	withoutParenthesis := tuple.Fn3(func(_ Character, x R, _ Character) R { return x })
	return ps.Map(
		withoutParenthesis,
		ps.Sequence3(OpenParenthesis, x, CloseParenthesis),
	)
}

func InParenthesesWhiteSpaceAllowed[R any](x func([]Character) tuple.Of2[rslt.Of[R], []Character]) func([]Character) tuple.Of2[rslt.Of[R], []Character] {
	withoutParenthesis := func(x tuple.Of3[Character, R, Character]) R {
		return x.X2()
	}
	return ps.Map(
		withoutParenthesis,
		IgnoreWhiteSpaceBetween3(OpenParenthesis, x, CloseParenthesis),
	)
}

func InBracesWhiteSpacesAllowed[R any](
	x func([]Character) tuple.Of2[rslt.Of[R], []Character],
) func([]Character) tuple.Of2[rslt.Of[R], []Character] {
	withoutBrace := func(x tuple.Of3[Character, R, Character]) R {
		return x.X2()
	}
	return ps.Map(
		withoutBrace,
		IgnoreWhiteSpaceBetween3(OpenBrace, x, CloseBrace),
	)
}

func InSquareBracketsWhiteSpacesAllowed[R any](
	x func([]Character) tuple.Of2[rslt.Of[R], []Character],
) func([]Character) tuple.Of2[rslt.Of[R], []Character] {
	withoutBrace := func(x tuple.Of3[Character, R, Character]) R {
		return x.X2()
	}
	return ps.Map(
		withoutBrace,
		IgnoreWhiteSpaceBetween3(OpenSquareBracket, x, CloseSquareBracket),
	)
}

func InSingleQuotes[R any](x func([]Character) tuple.Of2[rslt.Of[R], []Character]) func([]Character) tuple.Of2[rslt.Of[R], []Character] {
	withoutQuotes := func(x tuple.Of3[Character, R, Character]) R {
		return x.X2()
	}
	return ps.Map(
		withoutQuotes,
		ps.Sequence3(SingleQuote, x, SingleQuote),
	)
}

func InDoubleQuotes[R any](x func([]Character) tuple.Of2[rslt.Of[R], []Character]) func([]Character) tuple.Of2[rslt.Of[R], []Character] {
	withoutQuotes := func(x tuple.Of3[Character, R, Character]) R {
		return x.X2()
	}
	return ps.Map(
		withoutQuotes,
		ps.Sequence3(DoubleQuote, x, DoubleQuote),
	)
}

func IgnoreSpaceBetween2[R1, R2 any](
	p1 func([]Character) tuple.Of2[rslt.Of[R1], []Character],
	p2 func([]Character) tuple.Of2[rslt.Of[R2], []Character],
) func([]Character) tuple.Of2[rslt.Of[tuple.Of2[R1, R2]], []Character] {
	merge := tuple.Fn3(func(r1 R1, _ []Character, r2 R2) tuple.Of2[R1, R2] {
		return tuple.New2(r1, r2)
	})
	return ps.Map(
		merge,
		ps.Sequence3(p1, ps.OptionalGreedyRepeat(Space), p2),
	)
}

func IgnoreSpaceBetween3[R1, R2, R3 any](
	p1 func([]Character) tuple.Of2[rslt.Of[R1], []Character],
	p2 func([]Character) tuple.Of2[rslt.Of[R2], []Character],
	p3 func([]Character) tuple.Of2[rslt.Of[R3], []Character],
) func([]Character) tuple.Of2[rslt.Of[tuple.Of3[R1, R2, R3]], []Character] {
	merge := tuple.Fn2(func(r1 R1, r23 tuple.Of2[R2, R3]) tuple.Of3[R1, R2, R3] {
		return tuple.New3(r1, r23.X1(), r23.X2())
	})
	return ps.Map(
		merge,
		IgnoreSpaceBetween2(p1, IgnoreSpaceBetween2(p2, p3)),
	)
}

func IgnoreWhiteSpaceBetween2[R1, R2 any](
	p1 func([]Character) tuple.Of2[rslt.Of[R1], []Character],
	p2 func([]Character) tuple.Of2[rslt.Of[R2], []Character],
) func([]Character) tuple.Of2[rslt.Of[tuple.Of2[R1, R2]], []Character] {
	merge := tuple.Fn3(func(r1 R1, _ []Character, r2 R2) tuple.Of2[R1, R2] {
		return tuple.New2(r1, r2)
	})
	return ps.Map(merge, ps.Sequence3(p1, ps.OptionalGreedyRepeat(WhiteSpace), p2))
}

func IgnoreWhiteSpaceBetween3[R1, R2, R3 any](
	p1 func([]Character) tuple.Of2[rslt.Of[R1], []Character],
	p2 func([]Character) tuple.Of2[rslt.Of[R2], []Character],
	p3 func([]Character) tuple.Of2[rslt.Of[R3], []Character],
) func([]Character) tuple.Of2[rslt.Of[tuple.Of3[R1, R2, R3]], []Character] {
	merge := tuple.Fn2(func(r1 R1, r23 tuple.Of2[R2, R3]) tuple.Of3[R1, R2, R3] {
		return tuple.New3(r1, r23.X1(), r23.X2())
	})
	return ps.Map(
		merge,
		IgnoreWhiteSpaceBetween2(p1, IgnoreWhiteSpaceBetween2(p2, p3)),
	)
}

func SpaceSeparated2[R1, R2 any](
	p1 func([]Character) tuple.Of2[rslt.Of[R1], []Character],
	p2 func([]Character) tuple.Of2[rslt.Of[R2], []Character],
) func([]Character) tuple.Of2[rslt.Of[tuple.Of2[R1, R2]], []Character] {
	merge := tuple.Fn3(func(r1 R1, _ []Character, r2 R2) tuple.Of2[R1, R2] {
		return tuple.New2(r1, r2)
	})
	return ps.Map(
		merge,
		ps.Sequence3(p1, ps.GreedyRepeatAtLeastOnce(Space), p2),
	)
}

func SpaceSeparated3[R1, R2, R3 any](
	p1 func([]Character) tuple.Of2[rslt.Of[R1], []Character],
	p2 func([]Character) tuple.Of2[rslt.Of[R2], []Character],
	p3 func([]Character) tuple.Of2[rslt.Of[R3], []Character],
) func([]Character) tuple.Of2[rslt.Of[tuple.Of3[R1, R2, R3]], []Character] {
	merge := tuple.Fn5(func(r1 R1, _1 []Character, r2 R2, _2 []Character, r3 R3) tuple.Of3[R1, R2, R3] {
		return tuple.New3(r1, r2, r3)
	})
	return ps.Map(
		merge,
		ps.Sequence5(
			p1,
			ps.GreedyRepeatAtLeastOnce(Space),
			p2,
			ps.GreedyRepeatAtLeastOnce(Space),
			p3,
		),
	)
}

func WhiteSpaceSeparated2[R1, R2 any](
	p1 func([]Character) tuple.Of2[rslt.Of[R1], []Character],
	p2 func([]Character) tuple.Of2[rslt.Of[R2], []Character],
) func([]Character) tuple.Of2[rslt.Of[tuple.Of2[R1, R2]], []Character] {
	merge := tuple.Fn3(func(r1 R1, _ []Character, r2 R2) tuple.Of2[R1, R2] {
		return tuple.New2(r1, r2)
	})
	return ps.Map(
		merge,
		ps.Sequence3(p1, ps.GreedyRepeatAtLeastOnce(WhiteSpace), p2),
	)
}

func WhiteSpaceSeparated3[R1, R2, R3 any](
	p1 func([]Character) tuple.Of2[rslt.Of[R1], []Character],
	p2 func([]Character) tuple.Of2[rslt.Of[R2], []Character],
	p3 func([]Character) tuple.Of2[rslt.Of[R3], []Character],
) func([]Character) tuple.Of2[rslt.Of[tuple.Of3[R1, R2, R3]], []Character] {
	merge := tuple.Fn5(func(r1 R1, _1 []Character, r2 R2, _2 []Character, r3 R3) tuple.Of3[R1, R2, R3] {
		return tuple.New3(r1, r2, r3)
	})
	return ps.Map(
		merge,
		ps.Sequence5(
			p1,
			ps.GreedyRepeatAtLeastOnce(WhiteSpace),
			p2,
			ps.GreedyRepeatAtLeastOnce(WhiteSpace),
			p3,
		),
	)
}

func GreedyRepeatAtLeastOnceSpaceSeparated[R any](
	p func([]Character) tuple.Of2[rslt.Of[R], []Character],
) func([]Character) tuple.Of2[rslt.Of[[]R], []Character] {
	merge := tuple.Fn2(func(x R, xs []tuple.Of2[[]Character, R]) []R {
		return append([]R{x}, slc.Map(tuple.Of2ToX2[[]Character, R])(xs)...)
	})
	return ps.Map(merge, ps.Sequence2(
		p, ps.OptionalGreedyRepeat(ps.Sequence2(ps.GreedyRepeatAtLeastOnce(Space), p)),
	))
}

func GreedyRepeatAtLeastOnceWhiteSpaceSeparated[R any](
	p func([]Character) tuple.Of2[rslt.Of[R], []Character],
) func([]Character) tuple.Of2[rslt.Of[[]R], []Character] {
	merge := tuple.Fn2(func(x R, xs []tuple.Of2[[]Character, R]) []R {
		return append([]R{x}, slc.Map(tuple.Of2ToX2[[]Character, R])(xs)...)
	})
	return ps.Map(merge, ps.Sequence2(
		p,
		ps.OptionalGreedyRepeat(ps.Sequence2(ps.GreedyRepeatAtLeastOnce(WhiteSpace), p)),
	))
}

func OptionalGreedyRepeatSpaceSeparated[R any](
	p func([]Character) tuple.Of2[rslt.Of[R], []Character],
) func([]Character) tuple.Of2[rslt.Of[[]R], []Character] {
	return ps.First(
		GreedyRepeatAtLeastOnceSpaceSeparated(p),
		func(xs []Character) tuple.Of2[rslt.Of[[]R], []Character] {
			return tuple.New2(rslt.Value([]R{}), xs)
		},
	)
}

func OptionalGreedyRepeatWhiteSpaceSeparated[R any](
	p func([]Character) tuple.Of2[rslt.Of[R], []Character],
) func([]Character) tuple.Of2[rslt.Of[[]R], []Character] {
	return ps.First(
		GreedyRepeatAtLeastOnceWhiteSpaceSeparated(p),
		func(xs []Character) tuple.Of2[rslt.Of[[]R], []Character] {
			return tuple.New2(rslt.Value([]R{}), xs)
		},
	)
}

func GreedyRepeatAtLeastOnceIgnoreWhiteSpaceBetween[R any](
	p func([]Character) tuple.Of2[rslt.Of[R], []Character],
) func([]Character) tuple.Of2[rslt.Of[[]R], []Character] {
	withLeadingWhitespace := ps.Map(
		tuple.Of2ToX2[[]Character, R],
		ps.Sequence2(ps.OptionalGreedyRepeat(WhiteSpace), p),
	)
	rest := ps.OptionalGreedyRepeat(withLeadingWhitespace)
	merge := tuple.Fn2(func(x R, xs []R) []R { return append([]R{x}, xs...) })
	return ps.Map(merge, ps.Sequence2(p, rest))
}

func OptionalGreedyRepeatIgnoreWhiteSpaceBetween[R any](
	p func([]Character) tuple.Of2[rslt.Of[R], []Character],
) func([]Character) tuple.Of2[rslt.Of[[]R], []Character] {
	return ps.First(
		GreedyRepeatAtLeastOnceIgnoreWhiteSpaceBetween(p),
		func(xs []Character) tuple.Of2[rslt.Of[[]R], []Character] {
			return tuple.New2(rslt.Value([]R{}), xs)
		},
	)
}

func EndingWithCommaSpaceAllowed[R any](
	p func([]Character) tuple.Of2[rslt.Of[R], []Character],
) func([]Character) tuple.Of2[rslt.Of[R], []Character] {
	return ps.Map(
		tuple.Of2ToX1,
		IgnoreSpaceBetween2(
			p,
			Comma,
		),
	)
}

func OptionalLeadingWhiteSpace[R any](p func([]Character) tuple.Of2[rslt.Of[R], []Character]) func([]Character) tuple.Of2[rslt.Of[R], []Character] {
	return ps.Map(
		tuple.Of2ToX2[[]Character, R],
		ps.Sequence2(ps.OptionalGreedyRepeat(WhiteSpace), p),
	)
}

func GreedyRepeatedLinesInAutoIndentBlockAtLeastOnce[R any](
	line func([]Character) tuple.Of2[rslt.Of[R], []Character],
) func([]Character) tuple.Of2[rslt.Of[[]R], []Character] {
	errIndentRequired := errors.New("indent required")
	return func(xs []Character) tuple.Of2[rslt.Of[[]R], []Character] {
		firstIndent := ps.GreedyRepeatAtLeastOnce(Space)(xs)
		if ps.IsResultErr(firstIndent) {
			return tuple.New2(rslt.Error[[]R](errIndentRequired), xs)
		}
		indentWidth := len(ps.ResultValue(firstIndent))
		if indentWidth == 0 {
			return tuple.New2(rslt.Error[[]R](errIndentRequired), xs)
		}
		return GreedyRepeatedLinesInIndentBlockAtLeastOnce(line, indentWidth)(xs)
	}
}

func GreedyRepeatedLinesInIndentBlockAtLeastOnce[R any](
	line func([]Character) tuple.Of2[rslt.Of[R], []Character],
	indentWidth int,
) func([]Character) tuple.Of2[rslt.Of[[]R], []Character] {
	indentedLine := ps.Map(
		tuple.Of2ToX2[[]Character, R],
		ps.Sequence2(NTimesRepeat(Space, indentWidth), line),
	)
	indentedLineWithNewLine := ps.Map(
		tuple.Of2ToX2[[]Character, R],
		ps.Sequence2(ps.GreedyRepeatAtLeastOnce(LineBreak), indentedLine),
	)
	mergeMultiple := tuple.Fn2(func(p R, qs []R) []R { return append([]R{p}, qs...) })
	return ps.First(
		ps.Map(mergeMultiple, ps.Sequence2(
			indentedLine,
			ps.OptionalGreedyRepeat(indentedLineWithNewLine),
		)),
		ps.Map(slc.Pure, indentedLine),
	)
}

func NTimesRepeat[R any](
	p func([]Character) tuple.Of2[rslt.Of[R], []Character],
	n int,
) func([]Character) tuple.Of2[rslt.Of[[]R], []Character] {
	if n <= 0 {
		return func(xs []Character) tuple.Of2[rslt.Of[[]R], []Character] {
			return tuple.New2(rslt.Value([]R{}), xs)
		}
	}
	if n == 1 {
		return ps.Map(slc.Pure, p)
	}
	if n == 2 {
		return ps.Map(tuple.Fn2(func(a R, b R) []R { return []R{a, b} }), ps.Sequence2(p, p))
	}

	half := float64(n) / 2.
	l, r := int(math.Floor(half)), int(math.Ceil(half))
	concat := tuple.Fn2(func(a []R, b []R) []R { return append(slices.Clone(a), b...) })
	return ps.Map(concat, ps.Sequence2(
		NTimesRepeat[R](p, l),
		NTimesRepeat[R](p, r),
	))
}
