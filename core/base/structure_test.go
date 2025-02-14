package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
)

func TestStructure_Mutate(t *testing.T) {
	t.Run("should mutate value", func(t *testing.T) {
		result := NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewTag("hello"), NewString("world")),
			NewStructureRecord(NewTag("abc"), NewClass("f")),
		}).Mutate(newNormalMutationForTest(func(s string, object Object) optional.Of[Node] {
			if s == "f" {
				return optional.Value[Node](NewNumberFromString("123"))
			}
			return optional.Empty[Node]()
		})).Value()

		assert.Equal(t, NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewTag("hello"), NewString("world")),
			NewStructureRecord(NewTag("abc"), NewNumberFromString("123")),
		}), result)
	})

	t.Run("should not mutate if all values are already stable", func(t *testing.T) {
		mutation := newNormalMutationForTest(func(s string, object Object) optional.Of[Node] {
			if s == "f" {
				return optional.Value[Node](NewNumberFromString("123"))
			}
			return optional.Empty[Node]()
		})

		result := NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewTag("hello"), NewString("world")),
			NewStructureRecord(NewTag("abc"), NewClass("f")),
		}).Mutate(mutation). // stable after the first time
					Value().(Structure).
					Mutate(mutation)

		assert.True(t, result.IsEmpty())
	})
}

func TestStructure_MutateAsHead(t *testing.T) {
	t.Run("should mutate children first", func(t *testing.T) {
		result := NewStructureFromRecords([]StructureRecord{}).
			MutateAsHead(NewParamChain(slc.Pure([]Node{NewClass("f"), NewNumberFromString("456")})), newNormalMutationForTest(func(s string, object Object) optional.Of[Node] {
				if s == "f" {
					return optional.Value[Node](NewNumberFromString("123"))
				}
				return optional.Empty[Node]()
			})).Value()

		expected := NewOneLayerObject(NewStructureFromRecords([]StructureRecord{}), []Node{NewNumberFromString("123"), NewNumberFromString("456")})
		assert.Equal(t, expected, result)
	})

	t.Run("should process get tag", func(t *testing.T) {
		// result: {123: "hello"} (.get 123)
		result := NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewNumberFromString("123"), NewString("hello")),
		}).MutateAsHead(NewParamChain(slc.Pure([]Node{NewOneLayerObject(NewTag("get"), []Node{NewNumberFromString("123")})})), nil).Value()

		expected := NewString("hello")
		assert.Equal(t, expected, result)
	})

	t.Run("should not mutate when get with unknown key", func(t *testing.T) {
		result := NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewNumberFromString("123"), NewString("hello")),
		}).MutateAsHead(NewParamChain(slc.Pure([]Node{NewOneLayerObject(NewTag("get"), []Node{NewNumberFromString("999")})})), nil)

		assert.True(t, result.IsEmpty())
	})

	t.Run("should process set tag", func(t *testing.T) {
		result := NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewNumberFromString("123"), NewString("hello")),
		}).MutateAsHead(NewParamChain(slc.Pure([]Node{NewOneLayerObject(NewTag("set"), []Node{NewNumberFromString("123"), NewString("f")})})), nil).Value()

		expected := NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewNumberFromString("123"), NewString("f")),
		})
		assert.Equal(t, expected, result)
	})

	t.Run("should not mutate when set with unknown key", func(t *testing.T) {
		result := NewStructureFromRecords([]StructureRecord{
			NewStructureRecord(NewNumberFromString("123"), NewString("hello")),
		}).MutateAsHead(NewParamChain(slc.Pure([]Node{NewOneLayerObject(NewTag("set"), []Node{NewNumberFromString("999"), NewString("f")})})), nil)

		assert.True(t, result.IsEmpty())
	})
}
