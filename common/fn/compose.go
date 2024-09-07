package fn

func Compose[A1, A2, B any](f1 func(A2) B, f2 func(A1) A2) func(A1) B {
	return func(x A1) B {
		return f1(f2(x))
	}
}

func Compose3[A1, A2, A3, B any](f1 func(A3) B, f2 func(A2) A3, f3 func(A1) A2) func(A1) B {
	return func(x A1) B {
		return f1(f2(f3(x)))
	}
}

func Compose4[A1, A2, A3, A4, B any](f1 func(A4) B, f2 func(A3) A4, f3 func(A2) A3, f4 func(A1) A2) func(A1) B {
	return func(x A1) B {
		return f1(f2(f3(f4(x))))
	}
}
