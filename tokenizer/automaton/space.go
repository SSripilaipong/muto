package automaton

import "unicode"

type Space struct {
	namer
}

func newSpace() Space {
	return Space{
		namer: newConstNamer(NameSpace),
	}
}

func (s Space) Accept(x rune) (Abstract, bool) {
	if unicode.IsSpace(x) && x != '\n' {
		return newSpace(), true
	}
	return s, false
}
