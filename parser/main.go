package parser

import (
	"strings"

	"phi-lang/common/tuple"
	st "phi-lang/parser/syntaxtree"
	"phi-lang/parser/tokenizer"
)

var ParseToken = file

func ParseString(source string) []tuple.Of2[st.File, []tokenizer.Token] {
	tokens := TokensWithoutSpace(tokenizer.Tokenize(strings.NewReader(source)))
	return ParseToken(tokens)
}
