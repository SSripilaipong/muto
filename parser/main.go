package parser

import (
	"errors"
	"fmt"

	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree"
)

var ParseToken = parsing.Map(newPackage, file)

var ParseString = fn.Compose(ParseToken, psBase.StringToCharTokens)

func FilterResult(raw []tuple.Of2[syntaxtree.Package, []psBase.Character]) rslt.Of[syntaxtree.Package] {
	s := parsing.FilterSuccess(raw)
	if len(s) == 0 {
		var err error
		if len(raw) == 0 {
			err = errors.New("unknown parsing error")
		} else {
			c := raw[0].X2()[0]
			err = fmt.Errorf("parsing error at line %d, column %d: unexpected token '%c'", c.LineNumber(), c.ColumnNumber(), c.Value())
		}
		return rslt.Error[syntaxtree.Package](err)
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
