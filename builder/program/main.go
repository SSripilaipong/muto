package program

import (
	"github.com/SSripilaipong/go-common/rslt"

	"github.com/SSripilaipong/muto/builtin"
	"github.com/SSripilaipong/muto/builtin/global"
	"github.com/SSripilaipong/muto/builtin/portal"
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/core/module"
	fileParser "github.com/SSripilaipong/muto/parser/file"
	"github.com/SSripilaipong/muto/program"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

var BuildProgramFromString = fn.Compose(
	rslt.JoinFmap(BuildProgramFromSyntaxTree), fileParser.ParseModuleFromString,
)

func BuildProgramFromSyntaxTree(p st.Module) rslt.Of[program.Program] {
	baseMod := global.NewBaseModule()
	ports := portal.NewDefaultPortal()

	imported := builtin.NewBuiltinImportMapping(p.ImportNames()).Attach(ports)

	mod := module.BuildUserDefinedModuleFromBase(baseMod, p).
		Attach(module.NewDependency(ports, imported))
	return rslt.Value(program.New(mod))
}
