package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
	"github.com/SSripilaipong/muto/syntaxtree"
)

func TestString(t *testing.T) {
	r := String(StringToCharTokens(`"abc\n123\""abc`))
	expectedResult := syntaxtree.NewString(`"abc\n123\""`)
	expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
	assert.Equal(t, tuple.New2(rslt.Value(expectedResult), expectedRemainder), IgnoreLineAndColumnInNewResult(r))
}
