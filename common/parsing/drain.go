package parsing

import (
	"phi-lang/common/tuple"
)

func DrainLeading[S, R any](f func(S) bool, p func([]S) []tuple.Of2[R, []S]) func([]S) []tuple.Of2[R, []S] {
	return func(s []S) []tuple.Of2[R, []S] {
		i := 0
		for i < len(s) && f(s[i]) {
			i++
		}
		return p(s[i:])
	}
}

func DrainTrailing[S, R any](f func(S) bool, p func([]S) []tuple.Of2[R, []S]) func([]S) []tuple.Of2[R, []S] {
	return func(s []S) []tuple.Of2[R, []S] {
		var result []tuple.Of2[R, []S]
		for _, c := range p(s) {
			r, k := c.Return()
			if len(k) <= 0 {
				result = append(result, tuple.New2(r, k))
			} else {
				i := 0
				for i < len(k) && f(k[i]) {
					i++
				}
				result = append(result, tuple.New2(r, k[i:]))
			}
		}
		return result
	}
}
