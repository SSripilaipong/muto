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
	return fmt.Sprintf("token(%#v)", t.value)
}

type TokenType string

const (
	Character TokenType = "CHARACTER"
)

func IsCharacter(t Token) bool {
	return t.tokenType == Character
}

func TokenToValue(t Token) string {
	return t.Value()
}

func NewCharacter(x rune) Token {
	return NewToken(string(x), Character)
}
