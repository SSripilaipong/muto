package parsing

import "errors"

var ErrFirstNoMatch = errors.New("first: no parser matched")

func First[S, R any](ps ...Parser[R, S]) Parser[R, S] {
	if len(ps) == 0 {
		panic("ps must not be empty")
	}
	if len(ps) == 1 {
		return ps[0]
	}
	alternatives := First[S, R](ps[1:]...)
	return func(s []S) ParseResult[R, S] {
		first := ps[0](s)
		if first.IsOk() {
			return first
		}
		alt := alternatives(s)
		if alt.IsOk() {
			return alt
		}
		return NewParseResultError[R](ErrFirstNoMatch, s)
	}
}
