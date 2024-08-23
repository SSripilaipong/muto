package parser

import (
	"phi-lang/common/rslt"
	"phi-lang/parser/tokenizer"
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
