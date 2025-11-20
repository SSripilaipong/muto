package parsing

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
)

func Matches[S, R any](p func([]S) tuple.Of2[rslt.Of[R], []S]) func([]S) bool {
	return func(s []S) bool {
		return IsResultOk(p(s))
	}
}
