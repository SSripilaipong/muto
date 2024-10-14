package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/optional"
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
