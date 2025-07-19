package program

import (
	"errors"

	"github.com/SSripilaipong/muto/builtin/global"
	"github.com/SSripilaipong/muto/builtin/portal"
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/core/module"
	fileParser "github.com/SSripilaipong/muto/parser/file"
	"github.com/SSripilaipong/muto/program"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
)

var BuildProgramFromString = fn.Compose(
	rslt.JoinFmap(BuildProgramFromSyntaxTree), fileParser.ParsePackageFromString,
)

func BuildProgramFromSyntaxTree(p st.Package) rslt.Of[program.Program] {
	files := p.Files()
	if len(files) != 1 {
		return rslt.Error[program.Program](errors.New("currently only support exactly 1 file"))
	}

	mod := module.BuildModuleFromStatements(files[0].Statements()).
		Init(global.NewModule(), portal.NewDefaultPortal())
	return rslt.Value(program.New(mod))
}
