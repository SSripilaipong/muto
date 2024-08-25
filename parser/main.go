package parser

import (
	"errors"
	"fmt"
	"strings"

	"phi-lang/common/parsing"
	"phi-lang/common/rslt"
	"phi-lang/common/tuple"
	"phi-lang/parser/tokenizer"
	"phi-lang/syntaxtree"
)

var ParseToken = parsing.Map(newPackage, file)

func ParseString(source string) []tuple.Of2[syntaxtree.Package, []tokenizer.Token] {
	tokens := TokensWithoutSpace(tokenizer.Tokenize(strings.NewReader(source)))
	return ParseToken(tokens)
}

func ParseResult(s []tuple.Of2[syntaxtree.Package, []tokenizer.Token]) rslt.Of[syntaxtree.Package] {
	if len(s) == 0 {
		return rslt.Error[syntaxtree.Package](errors.New("unknown error"))
	}
	result, residue := s[0].Return()
	if len(residue) > 0 {
		return rslt.Error[syntaxtree.Package](fmt.Errorf("cannot parse at %v", residue[0]))
	}
	return rslt.Value(result)
}

func newPackage(f syntaxtree.File) syntaxtree.Package {
	return syntaxtree.NewPackage([]syntaxtree.File{f})
}
