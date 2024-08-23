package tokenizer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	var tokens []Token
	str := `(abc . def)("abc123\"\'\)");

	 	var _x_ <=  12.345;`
	tokenIter := Tokenize(strings.NewReader(str))
	for {
		if token, err := tokenIter().Return(); err != nil {
			break
		} else {
			tokens = append(tokens, token)
		}
	}

	assert.Equal(t, []Token{
		{value: "(", tokenType: Bracket},
		{value: "abc", tokenType: Identifier},
		{value: " ", tokenType: Space},
		{value: ".", tokenType: Symbol},
		{value: " ", tokenType: Space},
		{value: "def", tokenType: Identifier},
		{value: ")", tokenType: Bracket},
		{value: "(", tokenType: Bracket},
		{value: `"abc123\"\'\)"`, tokenType: String},
		{value: ")", tokenType: Bracket},
		{value: ";", tokenType: Symbol},
		{value: "\n\n", tokenType: LineBreak},
		{value: "\t \t", tokenType: Space},
		{value: "var", tokenType: Identifier},
		{value: " ", tokenType: Space},
		{value: "_x_", tokenType: Identifier},
		{value: " ", tokenType: Space},
		{value: "<=", tokenType: Symbol},
		{value: "  ", tokenType: Space},
		{value: "12.345", tokenType: Number},
		{value: ";", tokenType: Symbol},
	}, tokens)
}
