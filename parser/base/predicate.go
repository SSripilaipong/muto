package base

import (
	"slices"
	"unicode"
)

func IsBooleanValue(s string) bool {
	return s == "true" || s == "false"
}

func IsDoubleQuote(x rune) bool {
	return x == '"'
}

func IsSingleQuote(x rune) bool {
	return x == '\''
}

func IsDigit(x rune) bool {
	return unicode.IsDigit(x)
}

func IsIdentifierFirstLetter(x rune) bool {
	return unicode.IsLetter(x) || x == '_'
}

func IsIdentifierFirstLetterNonUpperCase(x rune) bool {
	return IsIdentifierFirstLetter(x) && !unicode.IsUpper(x) && x != '_'
}

func IsIdentifierFirstLetterUpperCaseAndUnderscore(x rune) bool {
	return IsIdentifierFirstLetterUpperCase(x) || x == '_'
}

func IsIdentifierFirstLetterUpperCase(x rune) bool {
	return IsIdentifierFirstLetter(x) && unicode.IsUpper(x)
}

func IsIdentifierFollowingLetter(x rune) bool {
	return unicode.IsLetter(x) || unicode.IsDigit(x) || slices.Contains([]rune{'_', '?', '\'', '!', '-'}, x)
}

func IsSymbol(x rune) bool {
	return (unicode.IsSymbol(x) || unicode.IsPunct(x)) && !isBracket(x) && x != '_'
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

func IsUnderscore(x rune) bool {
	return x == '_'
}

func IsHyphen(x rune) bool {
	return x == '-'
}

func IsASCIILetter(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}

func IsASCIIDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
