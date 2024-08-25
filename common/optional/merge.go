package optional

func MergeFn[A, T any](f func(A, T) Of[A]) func(Of[A], Of[T]) Of[A] {
	return func(a Of[A], x Of[T]) Of[A] {
		if !a.IsEmpty() && !x.IsEmpty() {
			return f(a.Value(), x.Value())
		}
		return Empty[A]()
	}
}
