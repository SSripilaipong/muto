package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/SSripilaipong/go-common/rslt"

	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree"
)

func TestCommand_importCommand(t *testing.T) {
	t.Run("should parse single path token", func(t *testing.T) {
		result := importCommand(psBase.StringToCharTokens(":import abc-12_3+xxx"))
		require.Len(t, result, 1)

		r := result[0]
		assert.Equal(t, rslt.Value(syntaxtree.NewImport([]string{"abc-12_3"})), r.X1())
		assert.Equal(t, "+xxx", psBase.CharactersToString(r.X2()))
	})

	t.Run("should parse multiple path token", func(t *testing.T) {
		result := importCommand(psBase.StringToCharTokens(":import abc/-12_3+xxx"))
		require.Len(t, result, 1)

		r := result[0]
		assert.Equal(t, rslt.Value(syntaxtree.NewImport([]string{"abc", "-12_3"})), r.X1())
		assert.Equal(t, "+xxx", psBase.CharactersToString(r.X2()))
	})
}
