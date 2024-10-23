package parser

import (
	"errors"
	"fmt"

	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

var ParseToken = ps.RsMap(newPackage, file)

var ParseString = fn.Compose(ParseToken, psBase.StringToCharTokens)

func FilterResult(raw []tuple.Of2[rslt.Of[st.Package], []psBase.Character]) rslt.Of[st.Package] {
	s := ps.FilterResult(raw)
	if len(s) == 0 {
		var err error
		if len(raw) == 0 {
			err = errors.New("unknown parsing error")
		} else {
			c := raw[0].X2()[0]
			err = fmt.Errorf("parsing error at line %d, column %d: unexpected token '%c'", c.LineNumber(), c.ColumnNumber(), c.Value())
		}
		return rslt.Error[st.Package](err)
	}
	r, k := s[0].Return()
	if len(k) > 0 {
		err := errors.New("unknown parsing error")
		if r.IsErr() {
			err = r.Error()
		}
		c := k[0]
		return rslt.Error[st.Package](fmt.Errorf("parsing error at line %d, column %d: %w", c.LineNumber(), c.ColumnNumber(), err))
	}
	if r.IsErr() {
		return rslt.Error[st.Package](fmt.Errorf("parsing error: %w", r.Error()))
	}
	return r
}

func newPackage(f st.File) st.Package {
	return st.NewPackage([]st.File{f})
}
