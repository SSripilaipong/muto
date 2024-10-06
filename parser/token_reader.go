package parser

import (
	"github.com/SSripilaipong/muto/common/rslt"
	tk "github.com/SSripilaipong/muto/parser/tokens"
)

func TokensWithoutSpace(iter func() rslt.Of[tk.Token]) []tk.Token {
	var tokens []tk.Token
	for {
		if token, err := iter().Return(); err != nil {
			break
		} else if !tk.IsSpace(token) {
			tokens = append(tokens, token)
		}
	}
	return tokens
}
