package parsing

import (
	"github.com/SSripilaipong/go-common/optional"
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
)

func GreedyOptional[S, R any](p func([]S) tuple.Of2[rslt.Of[R], []S]) func([]S) tuple.Of2[rslt.Of[optional.Of[R]], []S] {
	return First(
		Map(optional.Value[R], p),
		func(xs []S) tuple.Of2[rslt.Of[optional.Of[R]], []S] {
			return tuple.New2(rslt.Value(optional.Empty[R]()), xs)
		},
	)
}
