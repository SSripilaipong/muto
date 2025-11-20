package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
	"github.com/SSripilaipong/muto/syntaxtree"
)

func TestBoolean(t *testing.T) {
	t.Run("should parse true", func(t *testing.T) {
		r := Boolean(StringToCharTokens(`trueabc`))
		expectedResult := syntaxtree.NewBoolean(`true`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, tuple.New2(rslt.Value(expectedResult), expectedRemainder), IgnoreLineAndColumnInNewResult(r))
	})

	t.Run("should parse false", func(t *testing.T) {
		r := Boolean(StringToCharTokens(`falseabc`))
		expectedResult := syntaxtree.NewBoolean(`false`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, tuple.New2(rslt.Value(expectedResult), expectedRemainder), IgnoreLineAndColumnInNewResult(r))
	})
}
