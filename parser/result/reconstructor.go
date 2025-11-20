package result

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	psPattern "github.com/SSripilaipong/muto/parser/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func reconstructor() func(xs []psBase.Character) tuple.Of2[rslt.Of[stResult.Reconstructor], []psBase.Character] {
	backSlash := ps.Sequence2(
		psBase.BackSlash,
		ps.OptionalGreedyRepeat(psBase.WhiteSpace),
	)
	extractorPart := ps.Prefix(backSlash, psPattern.ParamPart())
	builderPart := ps.Map(
		castObjectResult,
		psBase.InSquareBracketsWhiteSpacesAllowed(nakedObjectMultilines),
	)
	toReconstructor := tuple.Fn2(stResult.NewReconstructor)
	return ps.Map(
		toReconstructor,
		psBase.IgnoreWhiteSpaceBetween2(extractorPart, builderPart),
	)
}
