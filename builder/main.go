package builder

import (
	"errors"

	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation"
	"github.com/SSripilaipong/muto/parser"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

var BuildFromString = fn.Compose3(
	rslt.JoinFmap(BuildFromSyntaxTree), parser.FilterResult, parser.ParseString,
)

func BuildFromSyntaxTree(p st.Package) rslt.Of[Program] {
	files := p.Files()
	if len(files) != 1 {
		return rslt.Error[Program](errors.New("currently only support exactly 1 file"))
	}
	return rslt.Value(Program{
		mutate: mutationFromFile(files[0]),
	})
}

func mutationFromFile(f st.File) func(x base.MutableNode) optional.Of[base.Node] {
	return mutation.NewFromStatements(f.Statements())
}
