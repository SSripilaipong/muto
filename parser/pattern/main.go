package pattern

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

var RsPattern = ps.Map(rslt.Value, Pattern())

func Pattern() func(xs []psBase.Character) []tuple.Of2[stPattern.NamedRule, []psBase.Character] {
	castWithParamPart := tuple.Fn2(func(class base.Class, params stPattern.ParamPart) stPattern.NamedRule {
		return stPattern.NewNamedRule(class.Name(), params)
	})
	castClass := func(class base.Class) stPattern.NamedRule {
		return stPattern.NewNamedRule(class.Name(), stPattern.ParamsToFixedParamPart([]base.PatternParam{}))
	}

	return ps.First(
		ps.Map(castWithParamPart, psBase.SpaceSeparated2(psBase.ClassRule, rulePatternParamPart())),
		ps.Map(castClass, psBase.ClassRule),
	)
}
