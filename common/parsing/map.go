package parsing

import "muto/common/tuple"

func Map[S, R1, R2 any](f func(R1) R2, p func([]S) []tuple.Of2[R1, []S]) func([]S) []tuple.Of2[R2, []S] {
	return func(s []S) []tuple.Of2[R2, []S] {
		var result []tuple.Of2[R2, []S]
		for _, c := range p(s) {
			r, k := c.Return()
			result = append(result, tuple.New2(f(r), k))
		}
		return result
	}
}
