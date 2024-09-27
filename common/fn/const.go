package fn

func Const[A, B any](y B) func(A) B {
	return func(A) B {
		return y
	}
}
