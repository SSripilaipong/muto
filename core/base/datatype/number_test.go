package datatype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {
	t.Run("should convert integer", func(t *testing.T) {
		assert.Equal(t, int32(123), NewNumber("123").intValue)
		assert.False(t, NewNumber("123").isFloat)
	})

	t.Run("should convert float", func(t *testing.T) {
		assert.Equal(t, 123.45, NewNumber("123.45").floatValue)
		assert.True(t, NewNumber("123.45").isFloat)
	})
}

func TestModInteger(t *testing.T) {
	t.Run("should mod integers", func(t *testing.T) {
		assert.Equal(t, NewNumber("2"), ModInteger(NewNumber("17"), NewNumber("3")).Value())
	})

	t.Run("should not mod floats", func(t *testing.T) {
		assert.True(t, ModInteger(NewNumber("17"), NewNumber("3.0")).IsEmpty())
		assert.True(t, ModInteger(NewNumber("17.0"), NewNumber("3")).IsEmpty())
		assert.True(t, ModInteger(NewNumber("17.0"), NewNumber("3.0")).IsEmpty())
	})

	t.Run("should not mod with zero", func(t *testing.T) {
		assert.True(t, ModInteger(NewNumber("17"), NewNumber("0")).IsEmpty())
		assert.True(t, ModInteger(NewNumber("17"), NewNumber("0.0")).IsEmpty())
	})
}

func TestDivInteger(t *testing.T) {
	t.Run("should div integers", func(t *testing.T) {
		assert.Equal(t, NewNumber("5"), DivInteger(NewNumber("17"), NewNumber("3")).Value())
		assert.Equal(t, NewNumber("-6"), DivInteger(NewNumber("-17"), NewNumber("3")).Value())
	})

	t.Run("should not div floats", func(t *testing.T) {
		assert.True(t, DivInteger(NewNumber("17"), NewNumber("3.0")).IsEmpty())
		assert.True(t, DivInteger(NewNumber("17.0"), NewNumber("3")).IsEmpty())
		assert.True(t, DivInteger(NewNumber("17.0"), NewNumber("3.0")).IsEmpty())
	})

	t.Run("should not div with zero", func(t *testing.T) {
		assert.True(t, DivInteger(NewNumber("17"), NewNumber("0")).IsEmpty())
		assert.True(t, DivInteger(NewNumber("17"), NewNumber("0.0")).IsEmpty())
	})
}
