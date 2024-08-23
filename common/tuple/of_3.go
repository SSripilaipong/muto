package tuple

type Of3[T1, T2, T3 any] struct {
	x1 T1
	x2 T2
	x3 T3
}

func (t Of3[T1, T2, T3]) X1() T1 {
	return t.x1
}

func (t Of3[T1, T2, T3]) X2() T2 {
	return t.x2
}

func (t Of3[T1, T2, T3]) X3() T3 {
	return t.x3
}

func (t Of3[T1, T2, T3]) Return() (T1, T2, T3) {
	return t.x1, t.x2, t.x3
}

func New3[T1, T2, T3 any](x1 T1, x2 T2, x3 T3) Of3[T1, T2, T3] {
	return Of3[T1, T2, T3]{
		x1: x1,
		x2: x2,
		x3: x3,
	}
}

func Fn3[A1, A2, A3, B any](f func(A1, A2, A3) B) func(Of3[A1, A2, A3]) B {
	return func(t Of3[A1, A2, A3]) B {
		return f(t.X1(), t.X2(), t.X3())
	}
}
