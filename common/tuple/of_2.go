package tuple

import "fmt"

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

func (t Of2[T1, T2]) String() string {
	return fmt.Sprintf("(%s, %s)", fmt.Sprint(t.x1), fmt.Sprint(t.x2))
}

func New2[T1, T2 any](x1 T1, x2 T2) Of2[T1, T2] {
	return Of2[T1, T2]{
		x1: x1,
		x2: x2,
	}
}

func Fn2[A1, A2, B any](f func(A1, A2) B) func(Of2[A1, A2]) B {
	return func(t Of2[A1, A2]) B {
		return f(t.X1(), t.X2())
	}
}

func Of2ToX1[T1, T2 any](x Of2[T1, T2]) T1 {
	return x.X1()
}

func Of2ToX2[T1, T2 any](x Of2[T1, T2]) T2 {
	return x.X2()
}

func Of2MapX2[T1, T2, R any](f func(T2) R) func(x Of2[T1, T2]) Of2[T1, R] {
	return func(x Of2[T1, T2]) Of2[T1, R] {
		return New2(x.X1(), f(x.X2()))
	}
}
