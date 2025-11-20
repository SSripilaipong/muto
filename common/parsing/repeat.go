package parsing

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	"github.com/SSripilaipong/muto/common/slc"
)

func OptionalGreedyRepeat[S, R any](p func([]S) tuple.Of2[rslt.Of[R], []S]) func([]S) tuple.Of2[rslt.Of[[]R], []S] {
	return func(s []S) tuple.Of2[rslt.Of[[]R], []S] {
		r := p(s)
		if IsResultErr(r) {
			return tuple.New2(rslt.Value([]R{}), s)
		}

		next := OptionalGreedyRepeat(p)(r.X2())
		if IsResultErr(next) {
			return tuple.New2(rslt.Error[[]R](ResultError(next)), s)
		}

		result := append(slc.Pure(ResultValue(r)), ResultValue(next)...)
		return tuple.New2(rslt.Value(result), next.X2())
	}
}

func GreedyRepeatAtLeastOnce[S, R any](p func([]S) tuple.Of2[rslt.Of[R], []S]) func([]S) tuple.Of2[rslt.Of[[]R], []S] {
	return func(s []S) tuple.Of2[rslt.Of[[]R], []S] {
		first := p(s)
		if IsResultErr(first) {
			return tuple.New2(rslt.Error[[]R](ResultError(first)), s)
		}

		rest := OptionalGreedyRepeat(p)(first.X2())
		result := append(slc.Pure(ResultValue(first)), ResultValue(rest)...)
		return tuple.New2(rslt.Value(result), rest.X2())
	}
}
