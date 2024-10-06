package tokens

import (
	"fmt"
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
	Number     TokenType = "NUMBER"
	String     TokenType = "STRING"
	Space      TokenType = "SPACE"
	Symbol     TokenType = "SYMBOL"
	Character  TokenType = "CHARACTER"
)

func IsSpace(t Token) bool {
	return t.tokenType == Space
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
