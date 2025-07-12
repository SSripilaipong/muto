package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func TestParamPart_fixedParam(t *testing.T) {
	rule := ParamPart()

	t.Run("should parse rune as param", func(t *testing.T) {
		r := rule(base.StringToCharTokens(`'x'`))
		expectedResult := stPattern.PatternsToParamPart([]stBase.Pattern{
			syntaxtree.NewRune(`'x'`),
		})
		assert.Equal(t, base.SingleResult(expectedResult, []base.Character{}), base.AsParserResult(r))
	})
}
