package parsing

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
)

// Prefix composes two tuple parsers, keeping the result from the second parser.
func Prefix[S, P, R any](
	prefix func([]S) tuple.Of2[rslt.Of[P], []S],
	p func([]S) tuple.Of2[rslt.Of[R], []S],
) func([]S) tuple.Of2[rslt.Of[R], []S] {
	ignorePrefix := tuple.Fn2(func(_ P, r R) R { return r })
	return Map(ignorePrefix, Sequence2(prefix, p))
}
