package parsing

import (
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/tuple"
)

func OptionalGreedyRepeat[S, R any](p func([]S) []tuple.Of2[R, []S]) func([]S) []tuple.Of2[[]R, []S] {
	return func(s []S) []tuple.Of2[[]R, []S] {
		var result []tuple.Of2[[]R, []S]
		for _, possibleCase := range p(s) {
			r, k1 := possibleCase.Return()
			repeat := OptionalGreedyRepeat[S, R](p)(k1)
			if len(repeat) == 0 {
				result = append(result, tuple.New2([]R{r}, k1))
			} else {
				for _, next := range repeat {
					rs, k2 := next.Return()
					result = append(result, tuple.New2(append([]R{r}, rs...), k2))
				}
			}
		}
		if len(result) == 0 {
			result = append(result, tuple.New2([]R{}, s))
		}
		return result
	}
}

func RsOptionalGreedyRepeat[S, R any](p func([]S) []tuple.Of2[rslt.Of[R], []S]) func([]S) []tuple.Of2[rslt.Of[[]R], []S] {
	return func(s []S) []tuple.Of2[rslt.Of[[]R], []S] {
		var result []tuple.Of2[rslt.Of[[]R], []S]
		for _, possibleCase := range p(s) {
			r, k1 := possibleCase.Return()
			if r.IsErr() {
				continue
			}

			repeat := RsOptionalGreedyRepeat[S, R](p)(k1)
			if len(repeat) == 0 {
				result = append(result, tuple.New2(rslt.Value(slc.Pure(r.Value())), k1))
			} else {
				for _, next := range repeat {
					rs, k2 := next.Return()
					if rs.IsErr() {
						result = append(result, tuple.New2(rslt.Value(slc.Pure(r.Value())), k1))
						continue
					}
					result = append(result, tuple.New2(rslt.Value(append(slc.Pure(r.Value()), rs.Value()...)), k2))
				}
			}
		}
		if len(result) == 0 {
			result = append(result, tuple.New2(rslt.Value([]R{}), s))
		}
		return result
	}
}

func GreedyRepeatAtLeastOnce[S, R any](p func([]S) []tuple.Of2[R, []S]) func([]S) []tuple.Of2[[]R, []S] {
	return func(s []S) []tuple.Of2[[]R, []S] {
		var result []tuple.Of2[[]R, []S]
		for _, possibleCase := range p(s) {
			r, k1 := possibleCase.Return()
			repeat := OptionalGreedyRepeat[S, R](p)(k1)
			if len(repeat) == 0 {
				result = append(result, tuple.New2([]R{r}, k1))
			} else {
				for _, next := range repeat {
					rs, k2 := next.Return()
					result = append(result, tuple.New2(append([]R{r}, rs...), k2))
				}
			}
		}
		return result
	}
}

func RsGreedyRepeatAtLeastOnce[S, R any](p func([]S) []tuple.Of2[rslt.Of[R], []S]) func([]S) []tuple.Of2[rslt.Of[[]R], []S] {
	merge := rslt.Fmap(tuple.Fn2(func(r R, rs []R) []R {
		return append(slc.Pure(r), rs...)
	}))
	return Map(merge, RsSequence2(p, RsOptionalGreedyRepeat(p)))
}
