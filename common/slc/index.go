package slc

func ValidIndex[T any](xs []T, i int) bool {
	return 0 <= i && i < len(xs)
}

func ValidSliceUntilIndex(i int) bool {
	return i >= 0
}

func ValidSliceFromIndex[T any](xs []T, i int) bool {
	return 0 <= i && i < len(xs)
}

func LastIndex[T any](xs []T) int {
	return len(xs) - 1
}
