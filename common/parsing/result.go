package parsing

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	"github.com/SSripilaipong/muto/common/slc"
)

func FilterSuccess[R, S any](rs []tuple.Of2[R, []S]) (ss []tuple.Of2[R, []S]) {
	for _, r := range rs {
		if len(r.X2()) == 0 {
			ss = append(ss, r)
		}
	}
	return
}

func FilterResult[R, S any](rs []tuple.Of2[rslt.Of[R], []S]) []tuple.Of2[rslt.Of[R], []S] {
	if len(rs) == 0 {
		return nil
	}

	var ss []tuple.Of2[rslt.Of[R], []S]
	farthestResult := rs[0]
	for _, r := range rs {
		if len(r.X2()) == 0 && r.X1().IsOk() {
			ss = append(ss, r)
		} else if len(r.X2()) < len(farthestResult.X2()) {
			farthestResult = r
		}
	}
	if len(ss) == 0 {
		return slc.Pure(farthestResult)
	}
	return ss
}

func Result[S, R any](r R) func([]S) []tuple.Of2[R, []S] {
	return func(s []S) []tuple.Of2[R, []S] {
		return slc.Pure(tuple.New2(r, s))
	}
}
