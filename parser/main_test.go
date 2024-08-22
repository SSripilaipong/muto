package parser

import (
	"fmt"
	"testing"

	"phi-lang/tokenizer"
)

func TestParser(t *testing.T) {
	parser := NewParser()
	tokens := []tokenizer.Token{
		tokenizer.NewToken("main", tokenizer.Identifier),
		tokenizer.NewToken("=", tokenizer.Symbol),
		tokenizer.NewToken("helloWorld", tokenizer.Identifier),
		tokenizer.NewToken("\\n", tokenizer.LineBreak),
	}
	fmt.Println(parser(tokens))
}
