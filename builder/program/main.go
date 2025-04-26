package program

import (
	"errors"

	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/core/mutation"
	"github.com/SSripilaipong/muto/core/mutation/builtin"
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
	return rslt.Value(program.New(mutation.NewFromStatements(files[0].Statements(), builtin.NewBuiltinMutatorsForStdio())))
}
