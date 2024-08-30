package slc

func Fold[A, T any](f func(A, T) A) func(zero A) func(xs []T) A {
	return func(zero A) func(xs []T) A {
		return func(xs []T) A {
			y := zero
			for _, x := range xs {
				y = f(y, x)
			}
			return y
		}
	}
}

func FoldGroup[A, T any, G comparable](merge func(A, T) A, group func(T) G) func(zero A) func(xs []T) []A {
	return func(zero A) func(xs []T) []A {
		return func(xs []T) []A {
			aggregate := make(map[G]A)
			for _, x := range xs {
				g := group(x)

				a, ok := aggregate[g]
				if !ok {
					a = zero
				}
				a = merge(a, x)
				aggregate[g] = a
			}

			values := make([]A, 0, len(aggregate))
			for _, v := range aggregate {
				values = append(values, v)
			}
			return values
		}
	}
}
