package parsing

func Map[S, R1, R2 any](f func(R1) R2, p Parser[R1, S]) Parser[R2, S] {
	return func(s []S) ParseResult[R2, S] {
		r := p(s)
		if r.IsError() {
			return NewParseResultError[R2](r.Error(), s)
		}
		return NewParseResultValue(f(r.Value()), r.Remaining())
	}
}
