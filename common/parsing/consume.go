package parsing

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/tuple"
)

func RsConsumeIf[S any](f func(S) bool) func([]S) []tuple.Of2[rslt.Of[S], []S] {
	return func(s []S) []tuple.Of2[rslt.Of[S], []S] {
		if len(s) <= 0 {
			return slc.Pure(tuple.New2(rslt.Error[S](fmt.Errorf("unknown error")), s))
		}
		e := s[0]
		if f(e) {
			return slc.Pure(tuple.New2(rslt.Value(e), s[1:]))
		}
		return slc.Pure(tuple.New2(rslt.Error[S](fmt.Errorf("unknown error")), s))
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

func Transform[S, R any](f func(x S) rslt.Of[R]) func([]S) []tuple.Of2[R, []S] {
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
