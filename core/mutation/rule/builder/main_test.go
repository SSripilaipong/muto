package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	"github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestBuildBoolean(t *testing.T) {
	factory := NewLiteralBuilderFactory(newCoreLiteralBuilderFactory(NewStaticClassCollection()))
	t.Run("should build true", func(t *testing.T) {
		assert.Equal(t, base.NewBoolean(true), factory.NewBuilder(syntaxtree.NewBoolean("true")).Build(nil).Value())
	})

	t.Run("should build false", func(t *testing.T) {
		assert.Equal(t, base.NewBoolean(false), factory.NewBuilder(syntaxtree.NewBoolean("false")).Build(nil).Value())
	})
}

func TestBuildTag(t *testing.T) {
	factory := NewLiteralBuilderFactory(newCoreLiteralBuilderFactory(NewStaticClassCollection()))
	t.Run("should build tag", func(t *testing.T) {
		assert.Equal(t, base.NewTag("abc"), factory.NewBuilder(syntaxtree.NewTag(".abc")).Build(nil).Value())
	})
}

func TestNew_Structure(t *testing.T) {
	factory := NewLiteralBuilderFactory(newCoreLiteralBuilderFactory(NewStaticClassCollection()))
	t.Run("should build structure", func(t *testing.T) {
		assert.Equal(t, base.NewStructureFromRecords(nil), factory.NewBuilder(stResult.NewStructure([]stResult.StructureRecord{})).Build(nil).Value())
	})
}

func TestNew_Object(t *testing.T) {
	t.Run("should build nested object with no params", func(t *testing.T) {
		factory := NewSimplifiedNodeBuilderFactory(NewStaticClassCollection())
		template := stResult.NewObject(syntaxtree.NewClass("f"), stResult.FixedParamPart{})
		expectedResult := base.NewNamedOneLayerObject("f", nil)
		assert.Equal(t, expectedResult, factory.NewBuilder(template).Build(parameter.New()).Value())
	})

	t.Run("should not carry remaining children to param object", func(t *testing.T) {
		factory := NewSimplifiedNodeBuilderFactory(NewStaticClassCollection())
		template := stResult.NewNakedObject(
			syntaxtree.NewClass("f"),
			stResult.FixedParamPart{syntaxtree.NewClass("a")},
		)
		param := parameter.New().SetRemainingParamChain(base.NewParamChain([][]base.Node{{base.NewString("xxx")}}))
		expectedResult := base.NewNamedOneLayerObject("f", []base.Node{base.NewUnlinkedClass("a"), base.NewString("xxx")})
		assert.Equal(t, expectedResult, factory.NewBuilder(template).Build(param).Value())
	})
}
