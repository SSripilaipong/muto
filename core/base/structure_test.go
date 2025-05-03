package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
)

func TestStructure_Mutate(t *testing.T) {
	t.Run("should mutate value", func(t *testing.T) {
		// subject: {.hello: "world", .abc: (f)}
		f := NewClass("f", newNormalMutationForTest(func(object Object) optional.Of[Node] {
			return optional.Value[Node](NewNumberFromString("123"))
		}))
		result := NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewTag("hello"), NewString("world")),
			NewStructureRecord(NewTag("abc"), WrapWithObject(f)),
		}).Mutate().Value()

		assert.Equal(t, NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewTag("hello"), NewString("world")),
			NewStructureRecord(NewTag("abc"), NewNumberFromString("123")),
		}), result)
	})

	t.Run("should not mutate if all values are already stable", func(t *testing.T) {
		f := NewClass("f", newNormalMutationForTest(func(object Object) optional.Of[Node] {
			return optional.Value[Node](NewNumberFromString("123"))
		}))

		// subject: {.hello: "world", .abc: (f)}
		result := NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewTag("hello"), NewString("world")),
			NewStructureRecord(NewTag("abc"), WrapWithObject(f)),
		}).
			Mutate(). // stable after the first time
			Value().(Structure).
			Mutate()

		assert.True(t, result.IsEmpty())
	})
}

func TestStructure_MutateAsHead(t *testing.T) {
	t.Run("should mutate children first", func(t *testing.T) {
		f := NewClass("f", newNormalMutationForTest(func(object Object) optional.Of[Node] {
			return optional.Value[Node](NewNumberFromString("123"))
		}))

		// subject: {} (f) 456
		result := NewStructureFromRecords([]StructureRecord{}).
			MutateAsHead(NewParamChain(slc.Pure([]Node{WrapWithObject(f), NewNumberFromString("456")}))).Value()

		expected := NewOneLayerObject(NewStructureFromRecords([]StructureRecord{}), []Node{NewNumberFromString("123"), NewNumberFromString("456")})
		assert.Equal(t, expected, result)
	})

	t.Run("should process get tag", func(t *testing.T) {
		// result: {123: "hello"} (.get 123)
		result := NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewNumberFromString("123"), NewString("hello")),
		}).MutateAsHead(NewParamChain(slc.Pure([]Node{NewOneLayerObject(NewTag("get"), []Node{NewNumberFromString("123")})}))).Value()

		expected := NewString("hello")
		assert.Equal(t, expected, result)
	})

	t.Run("should not mutate when get with unknown key", func(t *testing.T) {
		result := NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewNumberFromString("123"), NewString("hello")),
		}).MutateAsHead(NewParamChain(slc.Pure([]Node{NewOneLayerObject(NewTag("get"), []Node{NewNumberFromString("999")})})))

		assert.True(t, result.IsEmpty())
	})

	t.Run("should process set tag", func(t *testing.T) {
		result := NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewNumberFromString("123"), NewString("hello")),
		}).MutateAsHead(NewParamChain(slc.Pure([]Node{NewOneLayerObject(NewTag("set"), []Node{NewNumberFromString("123"), NewString("f")})}))).Value()

		expected := NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewNumberFromString("123"), NewString("f")),
		})
		assert.Equal(t, expected, result)
	})

	t.Run("should not mutate when set with unknown key", func(t *testing.T) {
		result := NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewNumberFromString("123"), NewString("hello")),
		}).MutateAsHead(NewParamChain(slc.Pure([]Node{NewOneLayerObject(NewTag("set"), []Node{NewNumberFromString("999"), NewString("f")})})))

		assert.True(t, result.IsEmpty())
	})
}
