package parsing

import (
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/tuple"
)

func ConsumeOne[S, R any](f func(x S) rslt.Of[R]) func([]S) []tuple.Of2[R, []S] {
	return func(s []S) []tuple.Of2[R, []S] {
		if len(s) <= 0 {
			return nil
		}
		e := f(s[0])
		if e.IsOk() {
			return []tuple.Of2[R, []S]{tuple.New2(e.Value(), s[1:])}
		}
		return nil
	}
}

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
