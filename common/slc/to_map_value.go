package slc

func ToMapValue[T any, K comparable](key func(T) K) func([]T) map[K]T {
	return func(xs []T) map[K]T {
		m := make(map[K]T)
		for _, x := range xs {
			m[key(x)] = x
		}
		return m
	}
}
