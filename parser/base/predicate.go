package base

import (
	"slices"
	"unicode"
)

func IsBooleanValue(s string) bool {
	return s == "true" || s == "false"
}

func IsAtSign(x rune) bool {
	return x == '@'
}

func IsEqualSign(x rune) bool {
	return x == '='
}

func IsOpenParenthesis(x rune) bool {
	return x == '('
}

func IsCloseParenthesis(x rune) bool {
	return x == ')'
}

func IsDoubleQuote(x rune) bool {
	return x == '"'
}

func IsBackSlash(x rune) bool {
	return x == '\\'
}

func IsDigit(x rune) bool {
	return unicode.IsDigit(x)
}

func IsMinusSign(x rune) bool {
	return x == '-'
}

func IsDot(x rune) bool {
	return x == '.'
}

func IsIdentifierFirstLetter(x rune) bool {
	return unicode.IsLetter(x) || x == '_'
}

func IsIdentifierFirstLetterLowerCase(x rune) bool {
	return IsIdentifierFirstLetter(x) && unicode.IsLower(x)
}

func IsIdentifierFirstLetterUpperCase(x rune) bool {
	return IsIdentifierFirstLetter(x) && unicode.IsUpper(x)
}

func IsIdentifierFollowingLetter(x rune) bool {
	return unicode.IsLetter(x) || unicode.IsDigit(x) || slices.Contains([]rune{'_', '?', '\'', '!', '-'}, x)
}

func IsSymbol(x rune) bool {
	return (unicode.IsSymbol(x) || unicode.IsPunct(x)) && !isBracket(x)
}

func isBracket(s rune) bool {
	return s == '{' || s == '}' || s == '[' || s == ']' || s == '(' || s == ')'
}

func IsSpace(x rune) bool {
	return x == ' ' || x == '\t'
}

func IsLineBreak(x rune) bool {
	return x == '\n'
}
