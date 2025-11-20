package parsing

import (
	"github.com/SSripilaipong/go-common/optional"
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	"github.com/SSripilaipong/muto/common/slc"
)

func FilterSuccess[R, S any](rs []tuple.Of2[R, []S]) (ss []tuple.Of2[R, []S]) {
	for _, r := range rs {
		if len(r.X2()) == 0 {
			ss = append(ss, r)
		}
	}
	return
}

// FilterResult collapses the non-deterministic parser result down to the
// canonical SingleResult semantics: prefer successful parses that consumed all
// tokens, and otherwise keep the branch that progressed the farthest so that
// callers can surface a meaningful error.
func FilterResult[R, S any](rs []tuple.Of2[rslt.Of[R], []S]) []tuple.Of2[rslt.Of[R], []S] {
	if len(rs) == 0 {
		return nil
	}

	var bestSuccess optional.Of[tuple.Of2[rslt.Of[R], []S]]
	farthestResult := rs[0]
	for _, r := range rs {
		switch {
		case IsResultOk(r):
			if bestSuccess.IsEmpty() || len(r.X2()) < len(bestSuccess.Value().X2()) {
				bestSuccess = optional.Value(r)
			}
		case len(r.X2()) < len(farthestResult.X2()):
			farthestResult = r
		}
	}
	if bestSuccess.IsNotEmpty() {
		return slc.Pure(bestSuccess.Value())
	}
	return slc.Pure(farthestResult)
}

func Result[S, R any](r R) func([]S) []tuple.Of2[R, []S] {
	return func(s []S) []tuple.Of2[R, []S] {
		return slc.Pure(tuple.New2(r, s))
	}
}

// IsResultOk reports whether the tuple contains a successful rslt value.
func IsResultOk[R, S any](r tuple.Of2[rslt.Of[R], []S]) bool {
	return r.X1().IsOk()
}

// IsResultErr reports whether the tuple contains an error result.
func IsResultErr[R, S any](r tuple.Of2[rslt.Of[R], []S]) bool {
	return r.X1().IsErr()
}

func ResultValue[R, S any](r tuple.Of2[rslt.Of[R], []S]) R {
	if r.X1().IsErr() {
		var zero R
		return zero
	}
	return r.X1().Value()
}

func ResultError[R, S any](r tuple.Of2[rslt.Of[R], []S]) error {
	if r.X1().IsOk() {
		return nil
	}
	return r.X1().Error()
}
