package file

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/go-common/rslt"

	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree"
)

func TestCommand_importCommand(t *testing.T) {
	t.Run("should parse single path token", func(t *testing.T) {
		result := importCommand(psBase.StringToCharTokens(":import abc-12_3+xxx"))

		assert.Equal(t, rslt.Value(syntaxtree.NewImport([]string{"abc-12_3"})), result.X1())
		assert.Equal(t, "+xxx", psBase.CharactersToString(result.X2()))
	})

	t.Run("should parse multiple path token", func(t *testing.T) {
		result := importCommand(psBase.StringToCharTokens(":import abc/-12_3+xxx"))

		assert.Equal(t, rslt.Value(syntaxtree.NewImport([]string{"abc", "-12_3"})), result.X1())
		assert.Equal(t, "+xxx", psBase.CharactersToString(result.X2()))
	})
}
