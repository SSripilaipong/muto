package global

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
)

func TestList_map(t *testing.T) {
	module := NewModule()
	class := module.GetClass("map")
	stringClass := module.GetClass("string")

	t.Run("should return empty list", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewOneLayerObject(class,
			stringClass, base.NewOneLayerObject(base.NewTag("my-data")),
		)) // map string ($)
		assert.Equal(t, base.NewOneLayerObject(base.NewUnlinkedRuleBasedClass("$")), result)
	})

	t.Run("should apply to each data", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewOneLayerObject(class,
			stringClass, base.NewOneLayerObject(base.NewTag("my-data"), base.NewNumberFromString("1"), base.NewNumberFromString("2")),
		)) // map string ($ 1 2)
		assert.Equal(t, base.NewConventionalList(base.NewString("1"), base.NewString("2")), result)
	})
}

func TestList_filter(t *testing.T) {
	module := NewModule()
	class := module.GetClass("filter")
	isStringClass := module.GetClass("string?")

	t.Run("should return empty list", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewOneLayerObject(class,
			isStringClass, base.NewOneLayerObject(base.NewTag("my-data")),
		)) // filter string? ($)
		assert.Equal(t, base.NewConventionalList(), result)
	})

	t.Run("should apply to each data", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewOneLayerObject(class,
			isStringClass, base.NewOneLayerObject(base.NewTag("my-data"),
				base.NewNumberFromString("1"), base.NewString("2"), base.NewNumberFromString("3"), base.NewString("4"),
			),
		)) // filter string? ($ 1 "2" 3 "4")
		assert.Equal(t, base.NewConventionalList(base.NewString("2"), base.NewString("4")), result)
	})
}

func TestList_parseRunesToString(t *testing.T) {
	resultStruct := func(result base.String, remainder base.Object) optional.Of[base.Node] {
		return optional.Value[base.Node](base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(base.ResultTag, result),
			base.NewStructureRecord(base.RemainderTag, remainder),
		}))
	}

	t.Run("should parse empty list", func(t *testing.T) {
		result := parseRunesToStringMutator.Mutate(base.NewNamedOneLayerObject(parseRunesToStringMutator.Name(),
			base.NewConventionalList(),
		))
		assert.Equal(t, resultStruct(base.NewString(""), base.NewConventionalList()), result)
	})

	t.Run("should parse list with non-rune", func(t *testing.T) {
		result := parseRunesToStringMutator.Mutate(base.NewNamedOneLayerObject(parseRunesToStringMutator.Name(),
			base.NewConventionalList(base.NewTag("x")),
		))
		assert.Equal(t, resultStruct(base.NewString(""), base.NewConventionalList(base.NewTag("x"))), result)
	})

	t.Run("should parse list with runes", func(t *testing.T) {
		result := parseRunesToStringMutator.Mutate(base.NewNamedOneLayerObject(parseRunesToStringMutator.Name(),
			base.NewConventionalList(base.NewRune('x'), base.NewRune('y')),
		))
		assert.Equal(t, resultStruct(base.NewString("xy"), base.NewConventionalList()), result)
	})

	t.Run("should parse list with runes and non-runes", func(t *testing.T) {
		result := parseRunesToStringMutator.Mutate(base.NewNamedOneLayerObject(parseRunesToStringMutator.Name(),
			base.NewConventionalList(base.NewRune('x'), base.NewRune('y'), base.NewNumberFromString("123"), base.NewBoolean(true)),
		))
		assert.Equal(t, resultStruct(base.NewString("xy"), base.NewConventionalList(base.NewNumberFromString("123"), base.NewBoolean(true))), result)
	})
}
