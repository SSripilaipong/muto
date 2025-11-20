package parsing

import (
	"errors"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
)

var (
	ErrConsumeEmptyInput    = errors.New("consume: empty input")
	ErrConsumePredicateFail = errors.New("consume: predicate failed")
)

func ConsumeIf[S any](f func(S) bool) func([]S) tuple.Of2[rslt.Of[S], []S] {
	return func(s []S) tuple.Of2[rslt.Of[S], []S] {
		if len(s) == 0 {
			return tuple.New2(rslt.Error[S](ErrConsumeEmptyInput), s)
		}
		e := s[0]
		if f(e) {
			return tuple.New2(rslt.Value(e), s[1:])
		}
		return tuple.New2(rslt.Error[S](ErrConsumePredicateFail), s)
	}
}

func Transform[S, R any](f func(x S) rslt.Of[R]) func([]S) tuple.Of2[rslt.Of[R], []S] {
	return func(s []S) tuple.Of2[rslt.Of[R], []S] {
		if len(s) == 0 {
			return tuple.New2(rslt.Error[R](ErrConsumeEmptyInput), s)
		}
		e := f(s[0])
		if e.IsOk() {
			return tuple.New2(e, s[1:])
		}
		return tuple.New2(e, s)
	}
}
