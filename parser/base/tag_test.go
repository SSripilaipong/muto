package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/syntaxtree"
)

func TestTag(t *testing.T) {
	t.Run("should parse tag like class with dot prefix", func(t *testing.T) {
		r := Tag(StringToCharTokens(`.a_bc-'!'.123`))
		expectedResult := syntaxtree.NewTag(`.a_bc-'!'`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`.123`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse class", func(t *testing.T) {
		r := Tag(StringToCharTokens(`abc`))
		assert.Equal(t, EmptyResult[syntaxtree.Tag](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}
