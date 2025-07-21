package file

import (
	"errors"
	"fmt"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree"
)

var ParseModuleFromString = fn.Compose(FilterResult, ParseModuleCombinationFromString)

var ParseModuleCombinationFromString = fn.Compose(module, psBase.StringToCharTokens)

var module = ps.RsMap(newModule, File)

func newModule(f syntaxtree.File) syntaxtree.Module {
	return syntaxtree.NewModule([]syntaxtree.File{f})
}

func FilterResult[T any](raw []tuple.Of2[rslt.Of[T], []psBase.Character]) rslt.Of[T] {
	s := ps.FilterResult(raw)
	if len(s) == 0 {
		var err error
		if len(raw) == 0 {
			err = errors.New("unknown parsing error")
		} else {
			c := raw[0].X2()[0]
			err = fmt.Errorf("parsing error at line %d, column %d: unexpected token '%c'", c.LineNumber(), c.ColumnNumber(), c.Value())
		}
		return rslt.Error[T](err)
	}
	r, k := s[0].Return()
	if len(k) > 0 {
		err := errors.New("unknown parsing error")
		if r.IsErr() {
			err = r.Error()
		}
		c := k[0]
		return rslt.Error[T](fmt.Errorf("parsing error at line %d, column %d: %w", c.LineNumber(), c.ColumnNumber(), err))
	}
	if r.IsErr() {
		return rslt.Error[T](fmt.Errorf("parsing error: %w", r.Error()))
	}
	return r
}
