package optional

func Map[A, B any](f func(A) B) func(Of[A]) Of[B] {
	return func(x Of[A]) Of[B] {
		if x.IsEmpty() {
			return Empty[B]()
		}
		return Value(f(x.Value()))
	}
}
