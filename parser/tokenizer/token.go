package tokenizer

import (
	"fmt"

	"github.com/SSripilaipong/muto/parser/tokenizer/automaton"
)

type Token struct {
	value     string
	tokenType TokenType
}

func NewToken(value string, tokenType TokenType) Token {
	return Token{
		value:     value,
		tokenType: tokenType,
	}
}

func (t Token) Value() string {
	return t.value
}

func (t Token) String() string {
	return fmt.Sprintf("%s(%#v)", t.tokenType, t.value)
}

type TokenType string

const (
	Identifier TokenType = "IDENTIFIER"
	Empty      TokenType = "EMPTY"
	Bracket    TokenType = "BRACKET"
	LineBreak  TokenType = "LINE_BREAK"
	Number     TokenType = "NUMBER"
	String     TokenType = "STRING"
	Space      TokenType = "SPACE"
	Symbol     TokenType = "SYMBOL"
	Character  TokenType = "CHARACTER"
	Unknown    TokenType = "UNKNOWN"
)

func automatonNameToTokenType(name automaton.Name) TokenType {
	switch name {
	case automaton.NameEmpty:
		return Empty
	case automaton.NameBracket:
		return Bracket
	case automaton.NameIdentifier:
		return Identifier
	case automaton.NameLineBreak:
		return LineBreak
	case automaton.NameNumber:
		return Number
	case automaton.NameString:
		return String
	case automaton.NameSpace:
		return Space
	case automaton.NameSymbol:
		return Symbol
	}
	return Unknown
}

func IsSpace(t Token) bool {
	return t.tokenType == Space
}

func IsLineBreak(t Token) bool {
	return t.tokenType == LineBreak
}

func IsIdentifier(t Token) bool {
	return t.tokenType == Identifier
}

func IsSymbol(t Token) bool {
	return t.tokenType == Symbol
}

func IsCharacter(t Token) bool {
	return t.tokenType == Character
}

func IsString(t Token) bool {
	return t.tokenType == String
}

func IsNumber(t Token) bool {
	return t.tokenType == Number
}

func TokenToValue(t Token) string {
	return t.Value()
}

func NewString(x string) Token {
	return NewToken(x, String)
}

func NewCharacter(x rune) Token {
	return NewToken(string(x), Character)
}

func NewNumber(x string) Token {
	return NewToken(x, Number)
}

func NewIdentifier(x string) Token {
	return NewToken(x, Identifier)
}

func NewSymbol(x string) Token {
	return NewToken(x, Symbol)
}
