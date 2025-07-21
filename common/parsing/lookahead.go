package parsing

import "github.com/SSripilaipong/go-common/tuple"

func Lookahead[S, R any](f func([]S) bool, p func([]S) []tuple.Of2[R, []S]) func([]S) []tuple.Of2[R, []S] {
	return func(s []S) (result []tuple.Of2[R, []S]) {
		for _, possibleCase := range p(s) {
			if f(possibleCase.X2()) {
				result = append(result, possibleCase)
			}
		}
		return
	}
}
