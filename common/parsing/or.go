package parsing

import "phi-lang/common/tuple"

func Or[S, R any](ps ...func([]S) []tuple.Of2[R, []S]) func([]S) []tuple.Of2[R, []S] {
	if len(ps) == 0 {
		panic("ps must not be empty")
	}
	if len(ps) == 1 {
		return ps[0]
	}
	alternatives := Or[S, R](ps[1:]...)
	return func(s []S) []tuple.Of2[R, []S] {
		return append(ps[0](s), alternatives(s)...)
	}
}
