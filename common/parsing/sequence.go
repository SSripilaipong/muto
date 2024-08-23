package parsing

import (
	"phi-lang/common/tuple"
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

func Sequence3[S, R1, R2, R3 any](
	p1 func([]S) []tuple.Of2[R1, []S],
	p2 func([]S) []tuple.Of2[R2, []S],
	p3 func([]S) []tuple.Of2[R3, []S],
) func([]S) []tuple.Of2[tuple.Of3[R1, R2, R3], []S] {
	return func(s []S) []tuple.Of2[tuple.Of3[R1, R2, R3], []S] {
		var result []tuple.Of2[tuple.Of3[R1, R2, R3], []S]
		for _, c1 := range p1(s) {
			r1, k1 := c1.Return()
			for _, c2 := range p2(k1) {
				r2, k2 := c2.Return()
				for _, c3 := range p3(k2) {
					r3, k3 := c3.Return()
					result = append(result, tuple.New2(tuple.New3(r1, r2, r3), k3))
				}
			}
		}
		return result
	}
}
