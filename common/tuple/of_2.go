package tuple

type Of2[T1, T2 any] struct {
	x1 T1
	x2 T2
}

func (t Of2[T1, T2]) X1() T1 {
	return t.x1
}

func (t Of2[T1, T2]) X2() T2 {
	return t.x2
}

func (t Of2[T1, T2]) Return() (T1, T2) {
	return t.x1, t.x2
}

func New2[T1, T2 any](x1 T1, x2 T2) Of2[T1, T2] {
	return Of2[T1, T2]{
		x1: x1,
		x2: x2,
	}
}
