package parser

import (
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/parser/tokenizer"
)

func TokensWithoutSpace(iter func() rslt.Of[tokenizer.Token]) []tokenizer.Token {
	var tokens []tokenizer.Token
	for {
		if token, err := iter().Return(); err != nil {
			break
		} else if !tokenizer.IsSpace(token) {
			tokens = append(tokens, token)
		}
	}
	return tokens
}
