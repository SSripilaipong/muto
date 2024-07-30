package automaton

type String struct {
	namer
}

func newString() String {
	return String{
		namer: newConstNamer(NameString),
	}
}

func (s String) Accept(x rune) (Abstract, bool) {
	if x == '"' {
		return newTerminatedString(), true
	} else if x == '\\' {
		return newEscapedString(), true
	}
	return newString(), true
}

type EscapedString struct {
	namer
}

func newEscapedString() EscapedString {
	return EscapedString{
		namer: newConstNamer(NameString),
	}
}

func (e EscapedString) Accept(x rune) (Abstract, bool) {
	return newString(), true
}

type TerminatedString struct {
	namer
}

func newTerminatedString() TerminatedString {
	return TerminatedString{
		namer: newConstNamer(NameString),
	}
}

func (s TerminatedString) Accept(rune) (Abstract, bool) {
	return s, false
}
