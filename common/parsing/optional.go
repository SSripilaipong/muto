package parsing

import "github.com/SSripilaipong/go-common/optional"

func GreedyOptional[S, R any](p Parser[R, S]) Parser[optional.Of[R], S] {
	return func(s []S) ParseResult[optional.Of[R], S] {
		r := p(s)
		if r.IsError() {
			return NewParseResultValue(optional.Empty[R](), s)
		}
		return NewParseResultValue(optional.Value(r.Value()), r.Remaining())
	}
}
