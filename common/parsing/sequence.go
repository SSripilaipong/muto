package parsing

import (
	"github.com/SSripilaipong/go-common/rslt"

	"github.com/SSripilaipong/go-common/tuple"
)

func Sequence2[S, R1, R2 any](
	p1 func([]S) tuple.Of2[rslt.Of[R1], []S],
	p2 func([]S) tuple.Of2[rslt.Of[R2], []S],
) func([]S) tuple.Of2[rslt.Of[tuple.Of2[R1, R2]], []S] {
	return func(s []S) tuple.Of2[rslt.Of[tuple.Of2[R1, R2]], []S] {
		r1 := p1(s)
		if IsResultErr(r1) {
			return tuple.New2(rslt.Error[tuple.Of2[R1, R2]](ResultError(r1)), s)
		}

		r2 := p2(r1.X2())
		if IsResultErr(r2) {
			return tuple.New2(rslt.Error[tuple.Of2[R1, R2]](ResultError(r2)), s)
		}

		return tuple.New2(rslt.Value(tuple.New2(ResultValue(r1), ResultValue(r2))), r2.X2())
	}
}

func Sequence3[S, R1, R2, R3 any](
	p1 func([]S) tuple.Of2[rslt.Of[R1], []S],
	p2 func([]S) tuple.Of2[rslt.Of[R2], []S],
	p3 func([]S) tuple.Of2[rslt.Of[R3], []S],
) func([]S) tuple.Of2[rslt.Of[tuple.Of3[R1, R2, R3]], []S] {
	return func(s []S) tuple.Of2[rslt.Of[tuple.Of3[R1, R2, R3]], []S] {
		r1 := p1(s)
		if IsResultErr(r1) {
			return tuple.New2(rslt.Error[tuple.Of3[R1, R2, R3]](ResultError(r1)), s)
		}

		r2 := p2(r1.X2())
		if IsResultErr(r2) {
			return tuple.New2(rslt.Error[tuple.Of3[R1, R2, R3]](ResultError(r2)), s)
		}

		r3 := p3(r2.X2())
		if IsResultErr(r3) {
			return tuple.New2(rslt.Error[tuple.Of3[R1, R2, R3]](ResultError(r3)), s)
		}

		return tuple.New2(
			rslt.Value(tuple.New3(ResultValue(r1), ResultValue(r2), ResultValue(r3))),
			r3.X2(),
		)
	}
}

func Sequence5[S, R1, R2, R3, R4, R5 any](
	p1 func([]S) tuple.Of2[rslt.Of[R1], []S],
	p2 func([]S) tuple.Of2[rslt.Of[R2], []S],
	p3 func([]S) tuple.Of2[rslt.Of[R3], []S],
	p4 func([]S) tuple.Of2[rslt.Of[R4], []S],
	p5 func([]S) tuple.Of2[rslt.Of[R5], []S],
) func([]S) tuple.Of2[rslt.Of[tuple.Of5[R1, R2, R3, R4, R5]], []S] {
	return func(s []S) tuple.Of2[rslt.Of[tuple.Of5[R1, R2, R3, R4, R5]], []S] {
		r1 := p1(s)
		if IsResultErr(r1) {
			return tuple.New2(rslt.Error[tuple.Of5[R1, R2, R3, R4, R5]](ResultError(r1)), s)
		}

		r2 := p2(r1.X2())
		if IsResultErr(r2) {
			return tuple.New2(rslt.Error[tuple.Of5[R1, R2, R3, R4, R5]](ResultError(r2)), s)
		}

		r3 := p3(r2.X2())
		if IsResultErr(r3) {
			return tuple.New2(rslt.Error[tuple.Of5[R1, R2, R3, R4, R5]](ResultError(r3)), s)
		}

		r4 := p4(r3.X2())
		if IsResultErr(r4) {
			return tuple.New2(rslt.Error[tuple.Of5[R1, R2, R3, R4, R5]](ResultError(r4)), s)
		}

		r5 := p5(r4.X2())
		if IsResultErr(r5) {
			return tuple.New2(rslt.Error[tuple.Of5[R1, R2, R3, R4, R5]](ResultError(r5)), s)
		}

		return tuple.New2(
			rslt.Value(tuple.New5(ResultValue(r1), ResultValue(r2), ResultValue(r3), ResultValue(r4), ResultValue(r5))),
			r5.X2(),
		)
	}
}
