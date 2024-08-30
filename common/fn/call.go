package fn

func CallWith[B, A any](x A) func(f func(A) B) B {
	return func(f func(A) B) B {
		return f(x)
	}
}
