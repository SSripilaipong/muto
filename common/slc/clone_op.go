package slc

import "slices"

func CloneAppend[T any, S []T](xs S, x T) []T {
	return append(slices.Clone(xs), x)
}

func CloneConcat[T any, S []T](xs, ys S) []T {
	return append(slices.Clone(xs), ys...)
}
