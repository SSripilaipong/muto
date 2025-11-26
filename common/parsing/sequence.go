package parsing

import "github.com/SSripilaipong/go-common/tuple"

func Sequence2[S, R1, R2 any](p1 Parser[R1, S], p2 Parser[R2, S]) Parser[tuple.Of2[R1, R2], S] {
	return func(s []S) ParseResult[tuple.Of2[R1, R2], S] {
		r1 := p1(s)
		if r1.IsError() {
			return NewParseResultError[tuple.Of2[R1, R2]](r1.Error(), s)
		}

		r2 := p2(r1.Remaining())
		if r2.IsError() {
			return NewParseResultError[tuple.Of2[R1, R2]](r2.Error(), s)
		}

		return NewParseResultValue(tuple.New2(r1.Value(), r2.Value()), r2.Remaining())
	}
}

func Sequence3[S, R1, R2, R3 any](
	p1 Parser[R1, S],
	p2 Parser[R2, S],
	p3 Parser[R3, S],
) Parser[tuple.Of3[R1, R2, R3], S] {
	return func(s []S) ParseResult[tuple.Of3[R1, R2, R3], S] {
		r1 := p1(s)
		if r1.IsError() {
			return NewParseResultError[tuple.Of3[R1, R2, R3]](r1.Error(), s)
		}

		r2 := p2(r1.Remaining())
		if r2.IsError() {
			return NewParseResultError[tuple.Of3[R1, R2, R3]](r2.Error(), s)
		}

		r3 := p3(r2.Remaining())
		if r3.IsError() {
			return NewParseResultError[tuple.Of3[R1, R2, R3]](r3.Error(), s)
		}

		return NewParseResultValue(
			tuple.New3(r1.Value(), r2.Value(), r3.Value()),
			r3.Remaining(),
		)
	}
}

func Sequence5[S, R1, R2, R3, R4, R5 any](
	p1 Parser[R1, S],
	p2 Parser[R2, S],
	p3 Parser[R3, S],
	p4 Parser[R4, S],
	p5 Parser[R5, S],
) Parser[tuple.Of5[R1, R2, R3, R4, R5], S] {
	return func(s []S) ParseResult[tuple.Of5[R1, R2, R3, R4, R5], S] {
		r1 := p1(s)
		if r1.IsError() {
			return NewParseResultError[tuple.Of5[R1, R2, R3, R4, R5]](r1.Error(), s)
		}

		r2 := p2(r1.Remaining())
		if r2.IsError() {
			return NewParseResultError[tuple.Of5[R1, R2, R3, R4, R5]](r2.Error(), s)
		}

		r3 := p3(r2.Remaining())
		if r3.IsError() {
			return NewParseResultError[tuple.Of5[R1, R2, R3, R4, R5]](r3.Error(), s)
		}

		r4 := p4(r3.Remaining())
		if r4.IsError() {
			return NewParseResultError[tuple.Of5[R1, R2, R3, R4, R5]](r4.Error(), s)
		}

		r5 := p5(r4.Remaining())
		if r5.IsError() {
			return NewParseResultError[tuple.Of5[R1, R2, R3, R4, R5]](r5.Error(), s)
		}

		return NewParseResultValue(
			tuple.New5(r1.Value(), r2.Value(), r3.Value(), r4.Value(), r5.Value()),
			r5.Remaining(),
		)
	}
}
