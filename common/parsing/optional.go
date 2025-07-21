package parsing

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/go-common/tuple"

	"github.com/SSripilaipong/muto/common/slc"
)

func GreedyOptional[S, R any](p func([]S) []tuple.Of2[R, []S]) func([]S) []tuple.Of2[optional.Of[R], []S] {
	return First(
		Map(optional.Value[R], p),
		func(xs []S) []tuple.Of2[optional.Of[R], []S] {
			return slc.Pure(tuple.New2(optional.Empty[R](), xs))
		},
	)
}
