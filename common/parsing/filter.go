package parsing

import "github.com/SSripilaipong/go-common/tuple"

func Filter[S, R any](f func(R) bool, p func([]S) []tuple.Of2[R, []S]) func([]S) []tuple.Of2[R, []S] {
	return func(s []S) []tuple.Of2[R, []S] {
		var result []tuple.Of2[R, []S]
		for _, c := range p(s) {
			if f(c.X1()) {
				result = append(result, c)
			}
		}
		return result
	}
}
