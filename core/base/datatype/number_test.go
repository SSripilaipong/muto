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
		assert.Equal(t, float32(123.45), NewNumber("123.45").floatValue)
		assert.True(t, NewNumber("123.45").isFloat)
	})
}
