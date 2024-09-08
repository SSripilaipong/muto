package parsing

import (
	"muto/common/tuple"
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
