package parsing

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
)

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

func First[S, R any](ps ...func([]S) []tuple.Of2[R, []S]) func([]S) []tuple.Of2[R, []S] {
	if len(ps) == 0 {
		panic("ps must not be empty")
	}
	if len(ps) == 1 {
		return ps[0]
	}
	alternatives := First[S, R](ps[1:]...)
	return func(s []S) []tuple.Of2[R, []S] {
		first := ps[0](s)
		if len(first) > 0 {
			return first
		}
		return alternatives(s)
	}
}

func RsFirst[S, R any](ps ...func([]S) []tuple.Of2[rslt.Of[R], []S]) func([]S) []tuple.Of2[rslt.Of[R], []S] {
	if len(ps) == 0 {
		panic("ps must not be empty")
	}
	if len(ps) == 1 {
		return ps[0]
	}
	alternatives := RsFirst[S, R](ps[1:]...)
	return func(s []S) []tuple.Of2[rslt.Of[R], []S] {
		first := CollapseMultiResult(ps[0](s))
		if r, ok := first.Return(); ok && r.X1().IsOk() {
			return ExpandSingleResult(r)
		}
		if r, ok := CollapseMultiResult(alternatives(s)).Return(); ok {
			return ExpandSingleResult(r)
		}
		return nil
	}
}
