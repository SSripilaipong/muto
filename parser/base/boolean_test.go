package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/syntaxtree"
)

func TestBoolean(t *testing.T) {
	t.Run("should parse true", func(t *testing.T) {
		r := Boolean(StringToCharTokens(`trueabc`))
		expectedResult := syntaxtree.NewBoolean(`true`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse false", func(t *testing.T) {
		r := Boolean(StringToCharTokens(`falseabc`))
		expectedResult := syntaxtree.NewBoolean(`false`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}
