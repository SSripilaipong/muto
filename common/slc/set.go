package slc

import (
	"maps"
	"slices"
)

func Union[T comparable, S []T](xs, ys S) S {
	set := make(map[T]struct{})
	for _, x := range xs {
		set[x] = struct{}{}
	}
	for _, y := range ys {
		set[y] = struct{}{}
	}
	return slices.Collect(maps.Keys(set))
}
