package extractor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func TestNonObjectFactory_TryNonObject(t *testing.T) {
	factory := newCorePatternFactory(nil).nonObject

	t.Run("should extract rune", func(t *testing.T) {
		extractor := factory.TryNonObject(st.NewRune("'x'")).Value()
		assert.True(t, extractor.Extract(base.NewRune('x')).IsNotEmpty())
	})

	t.Run("should extract imported class", func(t *testing.T) {
		extractor := factory.TryNonObject(st.NewImportedClass("mod", "f")).Value()
		assert.True(t, extractor.Extract(base.NewUnlinkedImportedClass("mod", "f")).IsNotEmpty())
	})
}
