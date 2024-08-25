package builder

import (
	"errors"

	"phi-lang/common/fn"
	"phi-lang/common/optional"
	"phi-lang/common/rslt"
	"phi-lang/core/base"
	"phi-lang/core/mutation"
	"phi-lang/parser"
	st "phi-lang/syntaxtree"
)

var BuildFromString = fn.Compose3(
	rslt.JoinFmap(BuildFromSyntaxTree), parser.ParseResult, parser.ParseString,
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

func mutationFromFile(f st.File) func(object base.Object) optional.Of[base.Node] {
	return mutation.NewFromStatements(f.Statements())
}
