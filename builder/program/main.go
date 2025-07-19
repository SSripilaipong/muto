package program

import (
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
	rslt.JoinFmap(BuildProgramFromSyntaxTree), fileParser.ParseModuleFromString,
)

func BuildProgramFromSyntaxTree(p st.Module) rslt.Of[program.Program] {
	mod := module.BuildModule(p).
		Init(global.NewModule(), portal.NewDefaultPortal())
	return rslt.Value(program.New(mod))
}
