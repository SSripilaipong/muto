package parsing

import (
	"errors"
)

var ErrFilterRejected = errors.New("result filtered out")

func Filter[S, R any](f func(R) bool, p Parser[R, S]) Parser[R, S] {
	return func(s []S) ParseResult[R, S] {
		r := p(s)
		if r.IsError() {
			return r
		}
		v := r.Value()
		if f(v) {
			return r
		}
		return NewParseResultError[R](ErrFilterRejected, s)
	}
}
