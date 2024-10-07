package slc

func Flatten[T any](xs [][]T) (y []T) {
	for _, x := range xs {
		y = append(y, x...)
	}
	return
}
