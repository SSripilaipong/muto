package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestBuildBoolean(t *testing.T) {
	t.Run("should build true", func(t *testing.T) {
		assert.Equal(t, base.NewBoolean(true), New(stBase.NewBoolean("true")).Build(nil).Value())
	})

	t.Run("should build false", func(t *testing.T) {
		assert.Equal(t, base.NewBoolean(false), New(stBase.NewBoolean("false")).Build(nil).Value())
	})
}

func TestBuildTag(t *testing.T) {
	t.Run("should build tag", func(t *testing.T) {
		assert.Equal(t, base.NewTag("abc"), New(stBase.NewTag(".abc")).Build(nil).Value())
	})
}

func TestNew_Structure(t *testing.T) {
	t.Run("should build structure", func(t *testing.T) {
		assert.Equal(t, base.NewStructureFromRecords(nil), New(stResult.NewStructure([]stResult.StructureRecord{})).Build(nil).Value())
	})
}

func TestNew_Object(t *testing.T) {
	t.Run("should build nested object with no params", func(t *testing.T) {
		template := stResult.NewObject(stBase.NewClass("f"), stResult.ParamsToFixedParamPart(stResult.FixedParamPart{}))
		expectedResult := base.NewNamedOneLayerObject("f", nil)
		assert.Equal(t, expectedResult, New(template).Build(nil).Value())
	})
}
