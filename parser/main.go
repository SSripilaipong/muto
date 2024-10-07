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

func FilterResult(s []tuple.Of2[syntaxtree.Package, []psBase.Character]) rslt.Of[syntaxtree.Package] {
	s = FilterSuccess(s)
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

func FilterSuccess(rs []tuple.Of2[syntaxtree.Package, []psBase.Character]) (ss []tuple.Of2[syntaxtree.Package, []psBase.Character]) {
	for _, r := range rs {
		if len(r.X2()) == 0 {
			ss = append(ss, r)
		}
	}
	return
}
