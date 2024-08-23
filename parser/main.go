package parser

import (
	"phi-lang/common/tuple"
	"phi-lang/parser/syntaxtree"
	"phi-lang/tokenizer"
)

func NewParser() func(s []tokenizer.Token) []tuple.Of2[syntaxtree.File, []tokenizer.Token] {
	return file
}
