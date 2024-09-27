package parser

import (
	"strings"
	"unicode"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
)

func isFirstLetterCapital(s string) bool {
	return unicode.IsUpper([]rune(s)[0])
}

func isKeyword(s string) bool {
	return isBooleanValue(s)
}

func isBooleanValue(s string) bool {
	return s == "true" || s == "false"
}

func noVarSuffix(name string) bool {
	return !hasSuffix3Dots(name)
}

func hasSuffix3Dots(s string) bool {
	return strings.HasSuffix(s, "...") && strings.Count(s, ".") == 3
}

func isAtSign(x tk.Token) bool {
	return strings.TrimSpace(x.Value()) == "@"
}

func isEqualSign(x tk.Token) bool {
	return strings.TrimSpace(x.Value()) == "="
}

func isOpenParenthesis(x tk.Token) bool {
	return strings.TrimSpace(x.Value()) == "("
}

func isCloseParenthesis(x tk.Token) bool {
	return strings.TrimSpace(x.Value()) == ")"
}

func withTrailingLineBreak[R any](p func([]tk.Token) []tuple.Of2[R, []tk.Token]) func([]tk.Token) []tuple.Of2[R, []tk.Token] {
	return ps.Map(
		tuple.Fn2(func(r R, _ tk.Token) R { return r }),
		ps.Sequence2(p, ps.ConsumeIf(tk.IsLineBreak)),
	)
}

func withLeadingLineBreak[R any](p func([]tk.Token) []tuple.Of2[R, []tk.Token]) func([]tk.Token) []tuple.Of2[R, []tk.Token] {
	return ps.Map(
		tuple.Fn2(func(_ tk.Token, r R) R { return r }),
		ps.Sequence2(ps.ConsumeIf(tk.IsLineBreak), p),
	)
}

func ignoreTrailingLineBreak[R any](p func([]tk.Token) []tuple.Of2[R, []tk.Token]) func([]tk.Token) []tuple.Of2[R, []tk.Token] {
	return ps.DrainTrailing(tk.IsLineBreak, p)
}

func ignoreLeadingLineBreak[R any](p func([]tk.Token) []tuple.Of2[R, []tk.Token]) func([]tk.Token) []tuple.Of2[R, []tk.Token] {
	return ps.DrainLeading(tk.IsLineBreak, p)
}

func inParentheses[T any](x func([]tk.Token) []tuple.Of2[T, []tk.Token]) func([]tk.Token) []tuple.Of2[T, []tk.Token] {
	return ps.Map(withoutParenthesis[T], ps.Sequence3(openParenthesis, x, closeParenthesis))
}

func withoutParenthesis[T any](x tuple.Of3[tk.Token, T, tk.Token]) T {
	return x.X2()
}
