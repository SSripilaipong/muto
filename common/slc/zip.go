package slc

func LeftZipApply[A, B any](fs []func(A) B) func(defaultValue B) func([]A) []B {
	return func(defaultValue B) func([]A) []B {
		return func(xs []A) (ys []B) {
			for i, f := range fs {
				y := defaultValue
				if ValidIndex(xs, i) {
					y = f(xs[i])
				}
				ys = append(ys, y)
			}
			return
		}
	}
}
