package parsing

import (
	"github.com/SSripilaipong/muto/common/tuple"
)

func Prefix[S, P, R any](
	prefix func([]S) []tuple.Of2[P, []S],
	p func([]S) []tuple.Of2[R, []S],
) func([]S) []tuple.Of2[R, []S] {
	ignorePrefix := tuple.Fn2(func(_ P, r R) R { return r })
	return Map(ignorePrefix, Sequence2(prefix, p))
}
