package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/syntaxtree"
)

func TestFixedVar(t *testing.T) {
	t.Run("should parse single letter variable", func(t *testing.T) {
		r := FixedVarWithUnderscore(StringToCharTokens(`X.123`))
		expectedResult := syntaxtree.NewVariable(`X`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`.123`))
		assert.Equal(t, tuple.New2(rslt.Value(expectedResult), expectedRemainder), IgnoreLineAndColumnInNewResult(r))
	})

	t.Run("should parse multiple letters variable", func(t *testing.T) {
		r := FixedVarWithUnderscore(StringToCharTokens(`XYabc-123!'?-1s.123`))
		expectedResult := syntaxtree.NewVariable(`XYabc-123!'?-1s`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`.123`))
		assert.Equal(t, tuple.New2(rslt.Value(expectedResult), expectedRemainder), IgnoreLineAndColumnInNewResult(r))
	})

	t.Run("should not parse identifier starting with small case", func(t *testing.T) {
		r := FixedVarWithUnderscore(StringToCharTokens(`xy`))
		assert.True(t, ps.IsResultErr(r))
	})

	t.Run("should not parse variadic var", func(t *testing.T) {
		r := FixedVarWithUnderscore(IgnoreLineAndColumn(StringToCharTokens(`X...`)))
		assert.True(t, ps.IsResultErr(r))
	})
}
