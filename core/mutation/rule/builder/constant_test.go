package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/syntaxtree"
)

func TestConstantBuilderFactory_NewBuilder(t *testing.T) {
	t.Run("should build rune", func(t *testing.T) {
		result := newConstantBuilderFactory().
			NewBuilder(syntaxtree.NewRune("'\\n'")).
			Value().
			Build(nil)
		assert.Equal(t, result.Value(), base.NewRune('\n'))
	})
}
