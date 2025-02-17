package pattern

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func nestedObjectRuleParamPattern() func(xs []psBase.Character) []tuple.Of2[base.PatternParam, []psBase.Character] {
	return func(xs []psBase.Character) []tuple.Of2[base.PatternParam, []psBase.Character] {
		return ps.Or(
			anonymousRulePattern(),
			ps.Map(stPattern.NamedRuleToParam, Pattern()),
			ps.Map(stPattern.VariableRulePatternToRulePatternParam, variableRulePattern()),
		)(xs)
	}
}
