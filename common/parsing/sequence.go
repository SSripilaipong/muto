package parsing

import (
	"errors"

	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/tuple"
)

func Sequence2[S, R1, R2 any](
	p1 func([]S) []tuple.Of2[R1, []S],
	p2 func([]S) []tuple.Of2[R2, []S],
) func([]S) []tuple.Of2[tuple.Of2[R1, R2], []S] {
	return func(s []S) []tuple.Of2[tuple.Of2[R1, R2], []S] {
		var result []tuple.Of2[tuple.Of2[R1, R2], []S]
		for _, c1 := range p1(s) {
			r1, k1 := c1.Return()
			for _, c2 := range p2(k1) {
				r2, k2 := c2.Return()
				result = append(result, tuple.New2(tuple.New2(r1, r2), k2))
			}
		}
		return result
	}
}

func RsSequence2[S, R1, R2 any](
	p1 func([]S) []tuple.Of2[rslt.Of[R1], []S],
	p2 func([]S) []tuple.Of2[rslt.Of[R2], []S],
) func([]S) []tuple.Of2[rslt.Of[tuple.Of2[R1, R2]], []S] {
	return func(s []S) []tuple.Of2[rslt.Of[tuple.Of2[R1, R2]], []S] {
		var result []tuple.Of2[rslt.Of[tuple.Of2[R1, R2]], []S]
		for _, c1 := range p1(s) {
			r1, k1 := c1.Return()
			if r1.IsErr() {
				result = append(result, tuple.New2(rslt.Error[tuple.Of2[R1, R2]](r1.Error()), k1))
				continue
			}
			for _, c2 := range p2(k1) {
				r2, k2 := c2.Return()
				if r2.IsErr() {
					result = append(result, tuple.New2(rslt.Error[tuple.Of2[R1, R2]](r2.Error()), k2))
					continue
				}
				result = append(result, tuple.New2(rslt.Value(tuple.New2(r1.Value(), r2.Value())), k2))
			}
		}
		if len(result) == 0 {
			result = append(result, tuple.New2(rslt.Error[tuple.Of2[R1, R2]](errors.New("unknown error")), s))
		}
		return result
	}
}

func RsSequence3[S, R1, R2, R3 any](
	p1 func([]S) []tuple.Of2[rslt.Of[R1], []S],
	p2 func([]S) []tuple.Of2[rslt.Of[R2], []S],
	p3 func([]S) []tuple.Of2[rslt.Of[R3], []S],
) func([]S) []tuple.Of2[rslt.Of[tuple.Of3[R1, R2, R3]], []S] {
	merge := tuple.Fn2(func(r1 R1, r23 tuple.Of2[R2, R3]) tuple.Of3[R1, R2, R3] {
		return tuple.New3(r1, r23.X1(), r23.X2())
	})
	return RsMap(merge, RsSequence2(p1, RsSequence2(p2, p3)))
}

func Sequence3[S, R1, R2, R3 any](
	p1 func([]S) []tuple.Of2[R1, []S],
	p2 func([]S) []tuple.Of2[R2, []S],
	p3 func([]S) []tuple.Of2[R3, []S],
) func([]S) []tuple.Of2[tuple.Of3[R1, R2, R3], []S] {
	merge := tuple.Fn2(func(r1 R1, r23 tuple.Of2[R2, R3]) tuple.Of3[R1, R2, R3] {
		return tuple.New3(r1, r23.X1(), r23.X2())
	})
	return Map(merge, Sequence2(p1, Sequence2(p2, p3)))
}

func Sequence5[S, R1, R2, R3, R4, R5 any](
	p1 func([]S) []tuple.Of2[R1, []S],
	p2 func([]S) []tuple.Of2[R2, []S],
	p3 func([]S) []tuple.Of2[R3, []S],
	p4 func([]S) []tuple.Of2[R4, []S],
	p5 func([]S) []tuple.Of2[R5, []S],
) func([]S) []tuple.Of2[tuple.Of5[R1, R2, R3, R4, R5], []S] {
	merge := tuple.Fn2(func(r12 tuple.Of2[R1, R2], r345 tuple.Of3[R3, R4, R5]) tuple.Of5[R1, R2, R3, R4, R5] {
		return tuple.New5(r12.X1(), r12.X2(), r345.X1(), r345.X2(), r345.X3())
	})
	return Map(merge, Sequence2(Sequence2(p1, p2), Sequence3(p3, p4, p5)))
}

func RsSequence5[S, R1, R2, R3, R4, R5 any](
	p1 func([]S) []tuple.Of2[rslt.Of[R1], []S],
	p2 func([]S) []tuple.Of2[rslt.Of[R2], []S],
	p3 func([]S) []tuple.Of2[rslt.Of[R3], []S],
	p4 func([]S) []tuple.Of2[rslt.Of[R4], []S],
	p5 func([]S) []tuple.Of2[rslt.Of[R5], []S],
) func([]S) []tuple.Of2[rslt.Of[tuple.Of5[R1, R2, R3, R4, R5]], []S] {
	merge := tuple.Fn2(func(r12 tuple.Of2[R1, R2], r345 tuple.Of3[R3, R4, R5]) tuple.Of5[R1, R2, R3, R4, R5] {
		return tuple.New5(r12.X1(), r12.X2(), r345.X1(), r345.X2(), r345.X3())
	})
	return RsMap(merge, RsSequence2(RsSequence2(p1, p2), RsSequence3(p3, p4, p5)))
}
