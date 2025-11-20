package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
	"github.com/SSripilaipong/muto/syntaxtree"
)

func TestNumber(t *testing.T) {
	t.Run("should parse positive number", func(t *testing.T) {
		r := Number(StringToCharTokens(`123abc`))
		expectedResult := syntaxtree.NewNumber(`123`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, tuple.New2(rslt.Value(expectedResult), expectedRemainder), IgnoreLineAndColumnInNewResult(r))
	})

	t.Run("should parse positive float number", func(t *testing.T) {
		r := Number(StringToCharTokens(`123.456abc`))
		expectedResult := syntaxtree.NewNumber(`123.456`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, tuple.New2(rslt.Value(expectedResult), expectedRemainder), IgnoreLineAndColumnInNewResult(r))
	})

	t.Run("should parse negative number", func(t *testing.T) {
		r := Number(StringToCharTokens(`-123abc`))
		expectedResult := syntaxtree.NewNumber(`-123`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, tuple.New2(rslt.Value(expectedResult), expectedRemainder), IgnoreLineAndColumnInNewResult(r))
	})
}
