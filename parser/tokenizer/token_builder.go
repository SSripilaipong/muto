package tokenizer

import (
	"phi-lang/common/tuple"
	automaton2 "phi-lang/parser/tokenizer/automaton"
)

type tokenBuilder struct {
	atm     automaton2.Abstract
	buffer  []rune
	residue []rune
}

func newTokenBuilder() *tokenBuilder {
	return &tokenBuilder{
		atm:     automaton2.New(),
		buffer:  nil,
		residue: nil,
	}
}

func (b *tokenBuilder) IsDone() bool {
	return len(b.residue) > 0
}

func (b *tokenBuilder) Consume(r rune) {
	if b.IsDone() {
		b.residue = append(b.residue, r)
		return
	}

	var accept bool
	b.atm, accept = b.atm.Accept(r)
	if !accept {
		b.residue = append(b.residue, r)
		return
	}
	b.buffer = append(b.buffer, r)
}

func (b *tokenBuilder) Build() tuple.Of2[Token, []rune] {
	return tuple.New2(NewToken(string(b.buffer), automatonNameToTokenType(b.atm.Name())), b.residue)
}
