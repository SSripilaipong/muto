package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func TestObject_head(t *testing.T) {
	rule := Object()

	t.Run("should parse rune as head", func(t *testing.T) {
		r := rule(base.StringToCharTokens(`('x' 1)`))
		assert.True(t, ps.IsResultOk(r))
		expectedResult := stPattern.NewNonDeterminantObject(
			syntaxtree.NewRune(`'x'`),
			stPattern.PatternsToParamPart([]stBase.Pattern{syntaxtree.NewNumber("1")}),
		)
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})
}
