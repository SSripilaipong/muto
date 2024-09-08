package parser

import (
	"errors"
	"fmt"
	"strings"

	"muto/common/parsing"
	"muto/common/rslt"
	"muto/common/tuple"
	"muto/parser/tokenizer"
	"muto/syntaxtree"
)

var ParseToken = parsing.Map(newPackage, file)

func ParseString(source string) []tuple.Of2[syntaxtree.Package, []tokenizer.Token] {
	tokens := TokensWithoutSpace(tokenizer.Tokenize(strings.NewReader(source)))
	return ParseToken(tokens)
}

func ParseResult(s []tuple.Of2[syntaxtree.Package, []tokenizer.Token]) rslt.Of[syntaxtree.Package] {
	//s = FilterSuccess(s)
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

func FilterSuccess(rs []tuple.Of2[syntaxtree.Package, []tokenizer.Token]) (ss []tuple.Of2[syntaxtree.Package, []tokenizer.Token]) {
	for _, r := range rs {
		if len(r.X2()) == 0 {
			ss = append(ss, r)
		}
	}
	return
}
