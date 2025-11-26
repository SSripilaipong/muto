package parsing

import "github.com/SSripilaipong/go-common/tuple"

// Prefix composes two parsers, keeping the result from the second parser.
func Prefix[S, P, R any](prefix Parser[P, S], p Parser[R, S]) Parser[R, S] {
	ignorePrefix := tuple.Fn2(func(_ P, r R) R { return r })
	return Map(ignorePrefix, Sequence2(prefix, p))
}
