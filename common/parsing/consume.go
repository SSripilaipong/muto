package parsing

import (
	"phi-lang/common/tuple"
)

func ConsumeIf[S any](f func(S) bool) func([]S) []tuple.Of2[S, []S] {
	return func(s []S) []tuple.Of2[S, []S] {
		if len(s) <= 0 {
			return nil
		}
		e := s[0]
		if f(e) {
			return []tuple.Of2[S, []S]{tuple.New2(e, s[1:])}
		}
		return nil
	}
}
