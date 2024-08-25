package slc

func Filter[T any](f func(T) bool) func([]T) []T {
	return func(xs []T) (ys []T) {
		for _, x := range xs {
			if f(x) {
				ys = append(ys, x)
			}
		}
		return
	}
}
