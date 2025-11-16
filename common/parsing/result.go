package parsing

import (
	"github.com/SSripilaipong/go-common/optional"
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	"github.com/SSripilaipong/muto/common/slc"
)

// SingleResult is the target representation for every parser combinator result.
// It contains the rslt.Of payload together with the remaining tokens that the
// parser left untouched. The surrounding refactor will update every parser to
// return this shape instead of []tuple.Of2[...].
type SingleResult[S, R any] = tuple.Of2[rslt.Of[R], []S]

// MultiResult captures the current non-deterministic representation so that
// helper functions can describe how to collapse it into SingleResult during the
// transition. Once all call sites are migrated this alias will disappear.
type MultiResult[S, R any] = []tuple.Of2[rslt.Of[R], []S]

// ExpandSingleResult temporarily wraps the target result representation back
// into the old multi-result slice so existing combinators remain unchanged
// while they are migrated one-by-one.
func ExpandSingleResult[R, S any](r SingleResult[S, R]) MultiResult[S, R] {
	return slc.Pure(r)
}

// CollapseMultiResult converts the legacy multi-branch results into the new
// SingleResult semantics using the same filtering logic that callers rely on
// today. The optional return indicates whether any branch existed at all so the
// caller can continue propagating "no parse" situations.
func CollapseMultiResult[R, S any](rs MultiResult[S, R]) optional.Of[SingleResult[S, R]] {
	filtered := FilterResult(rs)
	if len(filtered) == 0 {
		return optional.Empty[SingleResult[S, R]]()
	}
	return optional.Value(filtered[0])
}

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
		case r.X1().IsOk():
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
