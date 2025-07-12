package extractor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/syntaxtree"
)

func TestNonObjectFactory_TryNonObject(t *testing.T) {
	factory := newCorePatternFactory(nil).nonObject

	t.Run("should extract rune", func(t *testing.T) {
		extractor := factory.TryNonObject(syntaxtree.NewRune("'x'")).Value()
		assert.True(t, extractor.Extract(base.NewRune('x')).IsNotEmpty())
	})
}
