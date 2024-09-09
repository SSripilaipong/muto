package slc

func Pure[T any](x T) []T {
	return []T{x}
}
