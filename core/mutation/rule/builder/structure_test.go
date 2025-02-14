package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestBuildStructure(t *testing.T) {
	t.Run("should build empty structure", func(t *testing.T) {
		tree := stResult.NewStructure([]stResult.StructureRecord{})
		mutationData := parameter.New()
		expectedObject := base.NewStructureFromRecords(nil)
		assert.Equal(t, expectedObject, newStructureBuilder(tree).Build(mutationData).Value())
	})

	t.Run("should build string key and value", func(t *testing.T) {
		tree := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(stBase.NewString(`"abc"`), stBase.NewString(`"def"`)),
		})
		mutationData := parameter.New()
		expectedObject := base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(base.NewStringObject("abc"), base.NewStringObject("def")),
		})
		assert.Equal(t, expectedObject, newStructureBuilder(tree).Build(mutationData).Value())
	})

	t.Run("should build class key", func(t *testing.T) {
		tree := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(stBase.NewClass("f"), stBase.NewString(`"def"`)),
		})
		mutationData := parameter.New()
		expectedObject := base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(base.NewClassObject("f"), base.NewStringObject("def")),
		})
		assert.Equal(t, expectedObject, newStructureBuilder(tree).Build(mutationData).Value())
	})

	t.Run("should build variable value", func(t *testing.T) {
		tree := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(stBase.NewBoolean("true"), stBase.NewVariable("A")),
		})
		mutationData := parameter.New().
			WithVariableMappings(parameter.NewVariableMapping("A", base.NewOneLayerObject(base.NewClass("f"), []base.Node{base.NewNumberObjectFromString("123")}))).
			Value()
		expectedObject := base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(base.NewBooleanObject(true), base.NewOneLayerObject(base.NewClass("f"), []base.Node{base.NewNumberObjectFromString("123")})),
		})
		assert.Equal(t, expectedObject, newStructureBuilder(tree).Build(mutationData).Value())
	})

	t.Run("should build object value with variables", func(t *testing.T) {
		tree := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(
				stBase.NewTag(".e"),
				stResult.NewObject(
					stBase.NewVariable("A"),
					stResult.ParamsToFixedParamPart([]stResult.Param{stBase.NewNumber("123"), stBase.NewVariable("B")}),
				),
			),
		})
		mutationData := parameter.New().
			WithVariableMappings(parameter.NewVariableMapping("A", base.NewClassObject("f"))).Value().
			WithVariableMappings(parameter.NewVariableMapping("B", base.NewTagObject("t"))).Value()
		expectedObject := base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(
				base.NewTagObject("e"),
				base.NewOneLayerObject(
					base.NewClass("f"),
					[]base.Node{base.NewNumberObjectFromString("123"), base.NewTagObject("t")},
				),
			),
		})
		assert.Equal(t, expectedObject, newStructureBuilder(tree).Build(mutationData).Value())
	})
}
