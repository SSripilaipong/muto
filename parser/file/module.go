package file

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree"
)

var ParseModuleFromString = fn.Compose(psBase.FilterResult, ParseModuleCombinationFromString)

var ParseModuleCombinationFromString = fn.Compose(module, psBase.StringToCharTokens)

var module = ps.Map(newModule, ps.ToParser(File)).Legacy

func newModule(f syntaxtree.File) syntaxtree.Module {
	return syntaxtree.NewModule([]syntaxtree.File{f})
}
