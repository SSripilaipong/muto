package predicate

import (
	"slices"
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

func IsDoubleQuote(x string) bool {
	return strings.TrimSpace(x) == "\""
}

func IsBackSlash(x string) bool {
	return strings.TrimSpace(x) == "\\"
}

func IsFirstRuneDigit(x string) bool {
	return unicode.IsDigit([]rune(x)[0])
}

func IsMinusSign(x string) bool {
	return x == "-"
}

func IsDot(x string) bool {
	return x == "."
}

func IsIdentifierFirstLetter(x string) bool {
	c := []rune(x)[0]
	return unicode.IsLetter(c) || c == '_'
}

func IsIdentifierFirstLetterLowerCase(x string) bool {
	return IsIdentifierFirstLetter(x) && strings.ToLower(x) == x
}

func IsIdentifierFirstLetterUpperCase(x string) bool {
	return IsIdentifierFirstLetter(x) && strings.ToUpper(x) == x
}

func IsIdentifierFollowingLetter(s string) bool {
	x := []rune(s)[0]
	return unicode.IsLetter(x) || unicode.IsDigit(x) || slices.Contains([]rune{'_', '?', '\'', '!', '-'}, x)
}

func IsSymbolLetter(s string) bool {
	x := []rune(s)[0]
	return (unicode.IsSymbol(x) || unicode.IsPunct(x)) && !isBracket(x)
}

func isBracket(s rune) bool {
	return s == '{' || s == '}' || s == '[' || s == ']' || s == '(' || s == ')'
}
