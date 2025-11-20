package parsing

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
)

func Lookahead[S, R any](f func([]S) bool, p func([]S) tuple.Of2[rslt.Of[R], []S]) func([]S) tuple.Of2[rslt.Of[R], []S] {
	return func(s []S) tuple.Of2[rslt.Of[R], []S] {
		r := p(s)
		if IsResultErr(r) {
			return r
		}
		if f(r.X2()) {
			return r
		}
		return tuple.New2(rslt.Error[R](ErrNoParserResult), s)
	}
}
