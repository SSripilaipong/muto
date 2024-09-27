package predicate

import (
	"strings"
	"unicode"
)

func IsClassName(s string) bool {
	return !IsFirstLetterCapital(s) && !IsKeyword(s)
}

func IsFirstLetterCapital(s string) bool {
	return unicode.IsUpper([]rune(s)[0])
}

func IsKeyword(s string) bool {
	return IsBooleanValue(s)
}

func IsBooleanValue(s string) bool {
	return s == "true" || s == "false"
}

func IsVariableName(s string) bool {
	return IsFirstLetterCapital(s) && NoVarSuffix(s)
}

func IsVariadicVariable(s string) bool {
	return HasSuffix3Dots(s) && IsVariableName(s[:len(s)-3])
}

func NoVarSuffix(name string) bool {
	return !HasSuffix3Dots(name)
}

func HasSuffix3Dots(s string) bool {
	return strings.HasSuffix(s, "...") && strings.Count(s, ".") == 3
}

func IsAtSign(x string) bool {
	return strings.TrimSpace(x) == "@"
}

func IsEqualSign(x string) bool {
	return strings.TrimSpace(x) == "="
}

func IsSymbol(x string) bool {
	return !IsEqualSign(x) && !IsOpenParenthesis(x) && !IsCloseParenthesis(x)
}

func IsOpenParenthesis(x string) bool {
	return strings.TrimSpace(x) == "("
}

func IsCloseParenthesis(x string) bool {
	return strings.TrimSpace(x) == ")"
}
