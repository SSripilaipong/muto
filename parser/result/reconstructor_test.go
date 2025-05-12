package result

import (
	"testing"

	"github.com/stretchr/testify/assert"

	psBase "github.com/SSripilaipong/muto/parser/base"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestReconstructor(t *testing.T) {
	t.Run("should parse reconstructor with single variable and a class builder extractor", func(t *testing.T) {
		result := reconstructor()(psBase.StringToCharTokens(`\A [$]abc`))
		expectedResult := stResult.NewReconstructor(
			stPattern.PatternsToFixedParamPart([]stBase.Pattern{st.NewVariable("A")}),
			stResult.NewObject(st.NewClass("$"), stResult.FixedParamPart{}),
		)
		expectedRemainder := psBase.IgnoreLineAndColumn(psBase.StringToCharTokens("abc"))
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(psBase.IgnoreLineAndColumnInResult(result)))
	})
}
