package parsing

func OptionalGreedyRepeat[S, R any](p Parser[R, S]) Parser[[]R, S] {
	return func(s []S) ParseResult[[]R, S] {
		var result []R
		remaining := s
		for {
			r := p(remaining)
			if r.IsError() {
				return NewParseResultValue(result, remaining)
			}
			result = append(result, r.Value())
			remaining = r.Remaining()
		}
	}
}

func GreedyRepeatAtLeastOnce[S, R any](p Parser[R, S]) Parser[[]R, S] {
	repeat := OptionalGreedyRepeat(p)
	return func(s []S) ParseResult[[]R, S] {
		first := p(s)
		if first.IsError() {
			return NewParseResultError[[]R](first.Error(), s)
		}

		rest := repeat(first.Remaining())
		result := append([]R{first.Value()}, rest.Value()...)
		return NewParseResultValue(result, rest.Remaining())
	}
}
