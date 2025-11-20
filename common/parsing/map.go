package parsing

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
)

func Map[S, R1, R2 any](f func(R1) R2, p func([]S) tuple.Of2[rslt.Of[R1], []S]) func([]S) tuple.Of2[rslt.Of[R2], []S] {
	return func(s []S) tuple.Of2[rslt.Of[R2], []S] {
		r := p(s)
		if IsResultErr(r) {
			return tuple.New2(rslt.Error[R2](ResultError(r)), s)
		}
		return tuple.New2(rslt.Value(f(ResultValue(r))), r.X2())
	}
}

func DeRs[S, R any](f func([]S) []tuple.Of2[rslt.Of[R], []S]) func([]S) []tuple.Of2[R, []S] {
	return func(xs []S) []tuple.Of2[R, []S] {
		ys := f(xs)

		var result []tuple.Of2[R, []S]
		for _, y := range ys {
			r, k := y.Return()
			if r.IsOk() {
				result = append(result, tuple.New2(r.Value(), k))
			}
		}
		return result
	}
}
