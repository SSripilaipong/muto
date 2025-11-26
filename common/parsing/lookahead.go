package parsing

func Lookahead[S, R any](f func([]S) bool, p Parser[R, S]) Parser[R, S] {
	return func(s []S) ParseResult[R, S] {
		r := p(s)
		if r.IsError() {
			return r
		}
		if f(r.Remaining()) {
			return r
		}
		return NewParseResultError[R](ErrNoParserResult, s)
	}
}
