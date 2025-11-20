package parsing

import (
	"errors"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
)

var ErrFilterRejected = errors.New("result filtered out")

func Filter[S, R any](
	f func(R) bool,
	p func([]S) tuple.Of2[rslt.Of[R], []S],
) func([]S) tuple.Of2[rslt.Of[R], []S] {
	return func(s []S) tuple.Of2[rslt.Of[R], []S] {
		r := p(s)
		if IsResultErr(r) {
			return r
		}
		v := ResultValue(r)
		if f(v) {
			return r
		}
		return tuple.New2(rslt.Error[R](ErrFilterRejected), s)
	}
}
