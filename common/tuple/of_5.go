package tuple

type Of5[T1, T2, T3, T4, T5 any] struct {
	x1 T1
	x2 T2
	x3 T3
	x4 T4
	x5 T5
}

func (t Of5[T1, T2, T3, T4, T5]) X1() T1 {
	return t.x1
}

func (t Of5[T1, T2, T3, T4, T5]) X2() T2 {
	return t.x2
}

func (t Of5[T1, T2, T3, T4, T5]) X3() T3 {
	return t.x3
}

func (t Of5[T1, T2, T3, T4, T5]) X4() T4 {
	return t.x4
}

func (t Of5[T1, T2, T3, T4, T5]) X5() T5 {
	return t.x5
}

func (t Of5[T1, T2, T3, T4, T5]) Return() (T1, T2, T3, T4, T5) {
	return t.x1, t.x2, t.x3, t.x4, t.x5
}

func New5[T1, T2, T3, T4, T5 any](x1 T1, x2 T2, x3 T3, x4 T4, x5 T5) Of5[T1, T2, T3, T4, T5] {
	return Of5[T1, T2, T3, T4, T5]{
		x1: x1,
		x2: x2,
		x3: x3,
		x4: x4,
		x5: x5,
	}
}

func Fn5[A1, A2, A3, A4, A5, B any](f func(A1, A2, A3, A4, A5) B) func(Of5[A1, A2, A3, A4, A5]) B {
	return func(t Of5[A1, A2, A3, A4, A5]) B {
		return f(t.X1(), t.X2(), t.X3(), t.X4(), t.X5())
	}
}
