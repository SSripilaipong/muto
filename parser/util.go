package parser

import (
	"strings"
	"unicode"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	"github.com/SSripilaipong/muto/parser/tokenizer"
)

func isFirstLetterCapital(s string) bool {
	return unicode.IsUpper([]rune(s)[0])
}

func noVarSuffix(name string) bool {
	return !hasSuffix3Dots(name)
}

func hasSuffix3Dots(s string) bool {
	return strings.HasSuffix(s, "...") && strings.Count(s, ".") == 3
}

func isAtSign(x tokenizer.Token) bool {
	return strings.TrimSpace(x.Value()) == "@"
}

func isEqualSign(x tokenizer.Token) bool {
	return strings.TrimSpace(x.Value()) == "="
}

func isOpenParenthesis(x tokenizer.Token) bool {
	return strings.TrimSpace(x.Value()) == "("
}

func isCloseParenthesis(x tokenizer.Token) bool {
	return strings.TrimSpace(x.Value()) == ")"
}

func withTrailingLineBreak[R any](p func([]tokenizer.Token) []tuple.Of2[R, []tokenizer.Token]) func([]tokenizer.Token) []tuple.Of2[R, []tokenizer.Token] {
	return ps.Map(
		tuple.Fn2(func(r R, _ tokenizer.Token) R { return r }),
		ps.Sequence2(p, ps.ConsumeIf(tokenizer.IsLineBreak)),
	)
}

func withLeadingLineBreak[R any](p func([]tokenizer.Token) []tuple.Of2[R, []tokenizer.Token]) func([]tokenizer.Token) []tuple.Of2[R, []tokenizer.Token] {
	return ps.Map(
		tuple.Fn2(func(_ tokenizer.Token, r R) R { return r }),
		ps.Sequence2(ps.ConsumeIf(tokenizer.IsLineBreak), p),
	)
}

func ignoreTrailingLineBreak[R any](p func([]tokenizer.Token) []tuple.Of2[R, []tokenizer.Token]) func([]tokenizer.Token) []tuple.Of2[R, []tokenizer.Token] {
	return ps.DrainTrailing(tokenizer.IsLineBreak, p)
}

func ignoreLeadingLineBreak[R any](p func([]tokenizer.Token) []tuple.Of2[R, []tokenizer.Token]) func([]tokenizer.Token) []tuple.Of2[R, []tokenizer.Token] {
	return ps.DrainLeading(tokenizer.IsLineBreak, p)
}
