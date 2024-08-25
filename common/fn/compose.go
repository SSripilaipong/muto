package fn

func Compose[A1, A2, B any](f func(A2) B, g func(A1) A2) func(A1) B {
	return func(x A1) B {
		return f(g(x))
	}
}

func Compose3[A1, A2, A3, B any](f func(A3) B, g func(A2) A3, h func(A1) A2) func(A1) B {
	return func(x A1) B {
		return f(g(h(x)))
	}
}
