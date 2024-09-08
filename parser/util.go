package parser

import (
	"strings"
	"unicode"

	ps "muto/common/parsing"
	"muto/common/tuple"
	"muto/parser/tokenizer"
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

func isEqualSign(x tokenizer.Token) bool {
	return strings.TrimSpace(x.Value()) == "="
}

func isOpenParenthesis(x tokenizer.Token) bool {
	return strings.TrimSpace(x.Value()) == "("
}

func isCloseParenthesis(x tokenizer.Token) bool {
	return strings.TrimSpace(x.Value()) == ")"
}

func ignoreTrailingLineBreak[R any](p func([]tokenizer.Token) []tuple.Of2[R, []tokenizer.Token]) func([]tokenizer.Token) []tuple.Of2[R, []tokenizer.Token] {
	return ps.DrainTrailing(tokenizer.IsLineBreak, p)
}

func ignoreLeadingLineBreak[R any](p func([]tokenizer.Token) []tuple.Of2[R, []tokenizer.Token]) func([]tokenizer.Token) []tuple.Of2[R, []tokenizer.Token] {
	return ps.DrainLeading(tokenizer.IsLineBreak, p)
}
