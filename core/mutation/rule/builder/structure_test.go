package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	"github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestBuildStructure(t *testing.T) {
	builder := newStructureBuilderFactory(newCoreLiteralBuilderFactory())

	t.Run("should build empty structure", func(t *testing.T) {
		tree := stResult.NewStructure([]stResult.StructureRecord{})
		mutationData := parameter.New()
		expectedObject := base.NewStructureFromRecords(nil)
		assert.Equal(t, expectedObject, builder.NewBuilder(tree).Build(mutationData).Value())
	})

	t.Run("should build string key and value", func(t *testing.T) {
		tree := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(syntaxtree.NewString(`"abc"`), syntaxtree.NewString(`"def"`)),
		})
		mutationData := parameter.New()
		expectedObject := base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(base.NewString("abc"), base.NewString("def")),
		})
		assert.Equal(t, expectedObject, builder.NewBuilder(tree).Build(mutationData).Value())
	})

	t.Run("should build class key", func(t *testing.T) {
		tree := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(syntaxtree.NewClass("f"), syntaxtree.NewString(`"def"`)),
		})
		mutationData := parameter.New()
		expectedObject := base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(base.NewUnlinkedClass("f"), base.NewString("def")),
		})
		assert.Equal(t, expectedObject, builder.NewBuilder(tree).Build(mutationData).Value())
	})

	t.Run("should build variable value", func(t *testing.T) {
		tree := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(syntaxtree.NewBoolean("true"), syntaxtree.NewVariable("A")),
		})
		mutationData := parameter.New().
			WithVariableMapping(parameter.NewVariableMapping("A", base.NewOneLayerObject(base.NewUnlinkedClass("f"), []base.Node{base.NewNumberFromString("123")}))).
			Value()
		expectedObject := base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(base.NewBoolean(true), base.NewOneLayerObject(base.NewUnlinkedClass("f"), []base.Node{base.NewNumberFromString("123")})),
		})
		assert.Equal(t, expectedObject, builder.NewBuilder(tree).Build(mutationData).Value())
	})

	t.Run("should build object value with variables", func(t *testing.T) {
		tree := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(
				syntaxtree.NewTag(".e"),
				stResult.NewObject(
					syntaxtree.NewVariable("A"),
					stResult.ParamsToFixedParamPart([]stResult.Param{syntaxtree.NewNumber("123"), syntaxtree.NewVariable("B")}),
				),
			),
		})
		mutationData := parameter.New().
			WithVariableMapping(parameter.NewVariableMapping("A", base.NewUnlinkedClass("f"))).Value().
			WithVariableMapping(parameter.NewVariableMapping("B", base.NewTag("t"))).Value()
		expectedObject := base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(
				base.NewTag("e"),
				base.NewOneLayerObject(
					base.NewUnlinkedClass("f"),
					[]base.Node{base.NewNumberFromString("123"), base.NewTag("t")},
				),
			),
		})
		assert.Equal(t, expectedObject, builder.NewBuilder(tree).Build(mutationData).Value())
	})

	t.Run("should not carry remaining children to inner record object", func(t *testing.T) {
		// tree: {.k: .v}
		tree := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(syntaxtree.NewTag(".k"), syntaxtree.NewTag(".v")),
		})
		mutationData := parameter.New().
			SetRemainingParamChain(base.NewParamChain([][]base.Node{{base.NewTag("xxx")}}))
		expectedObject := base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(base.NewTag("k"), base.NewTag("v")),
		})
		assert.Equal(t, expectedObject, builder.NewBuilder(tree).Build(mutationData).Value())
	})
}
