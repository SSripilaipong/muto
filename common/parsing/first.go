package parsing

import (
	"errors"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
)

var ErrFirstNoMatch = errors.New("first: no parser matched")

func First[S, R any](ps ...func([]S) tuple.Of2[rslt.Of[R], []S]) func([]S) tuple.Of2[rslt.Of[R], []S] {
	if len(ps) == 0 {
		panic("ps must not be empty")
	}
	if len(ps) == 1 {
		return ps[0]
	}
	alternatives := First[S, R](ps[1:]...)
	return func(s []S) tuple.Of2[rslt.Of[R], []S] {
		first := ps[0](s)
		if IsResultOk(first) {
			return first
		}
		alt := alternatives(s)
		if IsResultOk(alt) {
			return alt
		}
		return tuple.New2(rslt.Error[R](ErrFirstNoMatch), s)
	}
}
