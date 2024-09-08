package automaton

import "unicode"

type Identifier struct {
	namer
}

func newIdentifier() Identifier {
	return Identifier{
		namer: newConstNamer(NameIdentifier),
	}
}

func (t Identifier) Accept(x rune) (Abstract, bool) {
	if unicode.IsLetter(x) || unicode.IsDigit(x) || x == '_' {
		return newIdentifier(), true
	}
	if x == '.' {
		return newIdentifierSuffixDots(), true
	}
	return t, false
}

type IdentifierSuffixDots struct {
	namer
}

func newIdentifierSuffixDots() IdentifierSuffixDots {
	return IdentifierSuffixDots{
		namer: newConstNamer(NameIdentifier),
	}
}

func (t IdentifierSuffixDots) Accept(x rune) (Abstract, bool) {
	return t, x == '.'
}
