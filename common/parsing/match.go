package parsing

func Matches[S, R any](p Parser[R, S]) func([]S) bool {
	return func(s []S) bool {
		return p(s).IsOk()
	}
}
