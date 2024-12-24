package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base/datatype"
)

func TestNodeEqual_Boolean(t *testing.T) {
	t.Run("should be equal with boolean", func(t *testing.T) {
		assert.True(t, NodeEqual(NewBoolean(true), NewBoolean(true)))
		assert.False(t, NodeEqual(NewBoolean(true), NewBoolean(false)))
		assert.False(t, NodeEqual(NewBoolean(false), NewBoolean(true)))
		assert.True(t, NodeEqual(NewBoolean(false), NewBoolean(false)))
	})

	t.Run("should not be equal with object", func(t *testing.T) {
		assert.False(t, NodeEqual(NewOneLayerObject(nil, nil), NewBoolean(true)))
		assert.False(t, NodeEqual(NewBoolean(true), NewOneLayerObject(nil, nil)))
	})

	t.Run("should not be equal with number", func(t *testing.T) {
		assert.False(t, NodeEqual(NewNumber(datatype.NewNumber("1")), NewBoolean(true)))
		assert.False(t, NodeEqual(NewBoolean(true), NewNumber(datatype.NewNumber("1"))))
	})

	t.Run("should not be equal with string", func(t *testing.T) {
		assert.False(t, NodeEqual(NewString("true"), NewBoolean(true)))
		assert.False(t, NodeEqual(NewBoolean(true), NewString("true")))
	})

	t.Run("should not be equal with class", func(t *testing.T) {
		assert.False(t, NodeEqual(NewClass("true"), NewBoolean(true)))
		assert.False(t, NodeEqual(NewBoolean(true), NewClass("true")))
	})
}

func TestNodeEqual_Tag(t *testing.T) {
	t.Run("should not be equal to other type", func(t *testing.T) {
		assert.False(t, NodeEqual(NewTag("a"), NewClass("a")))
	})

	t.Run("should not be equal to tag with different name", func(t *testing.T) {
		assert.False(t, NodeEqual(NewTag("a"), NewTag("b")))
	})

	t.Run("should be equal to tag with same name", func(t *testing.T) {
		assert.True(t, NodeEqual(NewTag("a"), NewTag("a")))
	})
}
