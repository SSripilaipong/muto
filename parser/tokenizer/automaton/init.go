package automaton

import (
	"unicode"
)

type Init struct {
	namer
}

func newInit() Init {
	return Init{
		namer: newConstNamer(NameEmpty),
	}
}

func (t Init) Accept(x rune) (Abstract, bool) {
	if unicode.IsLetter(x) || x == '_' {
		return newIdentifier(), true
	} else if unicode.IsDigit(x) {
		return newNumber(), true
	} else if x == '"' {
		return newString(), true
	} else if x == '\n' {
		return newLineBreak(), true
	} else if unicode.IsSpace(x) {
		return newSpace(), true
	} else if isBracket(x) {
		return newBracket(), true
	} else if x == '-' {
		return newLeadingNegative(), true
	} else if unicode.IsSymbol(x) || unicode.IsPunct(x) {
		return newSymbol(), true
	}
	return nil, false
}
