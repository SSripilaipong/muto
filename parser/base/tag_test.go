package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/syntaxtree"
)

func TestTag(t *testing.T) {
	t.Run("should parse tag like class with dot prefix", func(t *testing.T) {
		r := Tag(StringToCharTokens(`.a_bc-'!'.123`))
		expectedResult := syntaxtree.NewTag(`.a_bc-'!'`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`.123`))
		assert.Equal(t, tuple.New2(rslt.Value(expectedResult), expectedRemainder), IgnoreLineAndColumnInNewResult(r))
	})

	t.Run("should not parse class", func(t *testing.T) {
		r := Tag(StringToCharTokens(`abc`))
		assert.True(t, ps.IsResultErr(r))
	})
}
