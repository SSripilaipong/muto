package fn

func Not[T any](f func(x T) bool) func(T) bool {
	return func(x T) bool {
		return !f(x)
	}
}

func And[T any](f func(x T) bool, g func(x T) bool) func(T) bool {
	return func(x T) bool {
		return f(x) && g(x)
	}
}

func Or[T any](f func(x T) bool, g func(x T) bool) func(T) bool {
	return func(x T) bool {
		return f(x) || g(x)
	}
}
