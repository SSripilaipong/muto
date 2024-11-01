package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestBuildStructure(t *testing.T) {
	t.Run("should build empty structure", func(t *testing.T) {
		tree := stResult.NewStructure([]stResult.StructureRecord{})
		mutationData := parameter.New()
		expectedObject := base.NewStructureFromRecords(nil)
		assert.Equal(t, expectedObject, buildStructure(tree)(mutationData).Value())
	})

	t.Run("should build string key and value", func(t *testing.T) {
		tree := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(st.NewString(`"abc"`), st.NewString(`"def"`)),
		})
		mutationData := parameter.New()
		expectedObject := base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(base.NewString("abc"), base.NewString("def")),
		})
		assert.Equal(t, expectedObject, buildStructure(tree)(mutationData).Value())
	})

	t.Run("should build class key", func(t *testing.T) {
		tree := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(st.NewClass("f"), st.NewString(`"def"`)),
		})
		mutationData := parameter.New()
		expectedObject := base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(base.NewClass("f"), base.NewString("def")),
		})
		assert.Equal(t, expectedObject, buildStructure(tree)(mutationData).Value())
	})

	t.Run("should build variable value", func(t *testing.T) {
		tree := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(st.NewBoolean("true"), st.NewVariable("A")),
		})
		mutationData := parameter.New().
			WithVariableMappings(parameter.NewVariableMapping("A", base.NewObject(base.NewClass("f"), []base.Node{base.NewNumberFromString("123")}))).
			Value()
		expectedObject := base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(base.NewBoolean(true), base.NewObject(base.NewClass("f"), []base.Node{base.NewNumberFromString("123")})),
		})
		assert.Equal(t, expectedObject, buildStructure(tree)(mutationData).Value())
	})

	t.Run("should build object value with variables", func(t *testing.T) {
		tree := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(
				st.NewTag(".e"),
				stResult.NewObject(
					st.NewVariable("A"),
					stResult.ParamsToFixedParamPart([]stResult.Param{st.NewNumber("123"), st.NewVariable("B")}),
				),
			),
		})
		mutationData := parameter.New().
			WithVariableMappings(parameter.NewVariableMapping("A", base.NewClass("f"))).Value().
			WithVariableMappings(parameter.NewVariableMapping("B", base.NewTag("t"))).Value()
		expectedObject := base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(
				base.NewTag("e"),
				base.NewObject(
					base.NewClass("f"),
					[]base.Node{base.NewNumberFromString("123"), base.NewTag("t")},
				),
			),
		})
		assert.Equal(t, expectedObject, buildStructure(tree)(mutationData).Value())
	})
}
