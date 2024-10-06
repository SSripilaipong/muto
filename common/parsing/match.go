package parsing

import "github.com/SSripilaipong/muto/common/tuple"

func Matches[S, R any](f func([]S) []tuple.Of2[R, []S]) func([]S) bool {
	return func(s []S) bool {
		return len(f(s)) > 0
	}
}
