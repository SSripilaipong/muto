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
	return t, false
}
