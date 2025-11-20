package result

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func nestedNode() func([]psBase.Character) tuple.Of2[rslt.Of[stResult.Node], []psBase.Character] {
	return ps.First(
		psBase.InParenthesesWhiteSpaceAllowed(NakedObjectMultilines),
		nonObjectNestedNode(),
	)
}

func nonObjectNestedNode() func([]psBase.Character) tuple.Of2[rslt.Of[stResult.Node], []psBase.Character] {
	return ps.First(
		ps.Map(stResult.ToNode, structure),
		ps.Map(stResult.ToNode, reconstructor()),
	)
}
