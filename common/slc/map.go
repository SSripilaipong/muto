package slc

func Map[A, B any](f func(A) B) func([]A) []B {
	return func(xs []A) (ys []B) {
		for _, x := range xs {
			ys = append(ys, f(x))
		}
		return
	}
}

func Juxt[A, B any](fs []func(A) B) func(A) []B {
	return func(x A) (ys []B) {
		for _, f := range fs {
			ys = append(ys, f(x))
		}
		return
	}
}

func UniqueMap[A, B any, K comparable](key func(B) K, f func(A) B) func([]A) []B {
	return func(xs []A) (ys []B) {
		mem := make(map[K]struct{})
		for _, x := range xs {
			y := f(x)
			k := key(y)

			if _, exists := mem[k]; exists {
				continue
			}

			mem[k] = struct{}{}
			ys = append(ys, y)
		}
		return
	}
}
